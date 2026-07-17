// Package pubsubactivities is the message-level dedup store for Pub/Sub
// consumers. One row per (msg_id, service) tracks whether a delivery is being
// processed, has been processed, or failed — so a redelivered or
// re-published message is processed at-most-once in effect while failures
// remain retryable and auditable (last_error is kept).
package pubsubactivities

import (
	"context"
	"errors"
	"fmt"

	"github.com/global-torque/go-common/db/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// ErrInProgress tells push handlers to NACK a concurrent delivery. ACKing it
// would acknowledge the original delivery too, creating a crash-loss window.
var ErrInProgress = errors.New("pubsub activity is already processing")

// ErrClaimLost means a stale worker tried to complete a lease that has since
// been reclaimed. The caller must not overwrite the current owner's state.
var ErrClaimLost = errors.New("pubsub activity processing lease was lost")

// Status values mirror the pubsub_activity_status_t enum.
const (
	StatusProcessing = "processing"
	StatusProcessed  = "processed"
	StatusFailed     = "failed"

	// maxLastError caps stored error text so a pathological error string
	// can't bloat the row.
	maxLastError    = 4000
	processingLease = "10 minutes"
)

// RecordDomainDelivery stores the transport identity independently from the
// immutable outbox event id. One event can acquire a new Pub/Sub message id
// when operators republish from a DLQ, so neither value substitutes for the
// other in audit records.
func RecordDomainDelivery(
	ctx context.Context,
	repo db.Repository,
	eventID, service, pubsubMessageID string,
	attempt int,
) error {
	if pubsubMessageID == "" {
		return nil
	}
	_, err := repo.Exec(ctx, `
		INSERT INTO pubsub_domain_event_deliveries
		    (service, event_id, pubsub_message_id, delivery_attempt)
		VALUES ($1, $2::uuid, $3, $4)
		ON CONFLICT (service, pubsub_message_id) DO UPDATE
		SET delivery_attempt = GREATEST(
		        pubsub_domain_event_deliveries.delivery_attempt,
		        EXCLUDED.delivery_attempt
		    )
	`, service, eventID, pubsubMessageID, attempt)
	if err != nil {
		return fmt.Errorf("record domain event delivery (%s/%s): %w", service, pubsubMessageID, err)
	}
	return nil
}

// Claim atomically acquires processing ownership of (msgID, service).
//
// Returns true when the caller now owns the message: either it was never seen,
// or a previous attempt failed, or its processing lease expired after a worker
// crash. Returns false,nil only for a completed duplicate. A live concurrent
// processing claim returns ErrInProgress so the delivery is NACKed, never ACKed
// out from under a worker that might still crash.
func Claim(ctx context.Context, repo db.Repository, msgID, service, topic string, attempt int) (bool, error) {
	_, claimed, err := ClaimLease(ctx, repo, msgID, service, topic, attempt)
	return claimed, err
}

// ClaimLease is Claim with a fencing token. Scoped consumers must use the
// returned token for terminal updates so a worker that outlives its lease
// cannot overwrite a newer owner's result.
func ClaimLease(
	ctx context.Context,
	repo db.Repository,
	msgID, service, topic string,
	attempt int,
) (string, bool, error) {
	const q = `
INSERT INTO pubsub_activities (msg_id, service, status, topic, attempt, claim_token)
VALUES ($1, $2, 'processing', $3, $4, $6::uuid)
ON CONFLICT (msg_id, service) DO UPDATE
   SET status     = 'processing',
       attempt    = EXCLUDED.attempt,
       claim_token = EXCLUDED.claim_token,
       updated_at = now()
   WHERE pubsub_activities.status = 'failed'
      OR (pubsub_activities.status = 'processing'
          AND pubsub_activities.updated_at < now() - $5::interval)
RETURNING claim_token::text`

	claimToken := uuid.NewString()
	var gotToken string
	err := repo.QueryRow(ctx, q, msgID, service, topic, attempt, processingLease, claimToken).Scan(&gotToken)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			var status string
			if statusErr := repo.QueryRow(
				ctx,
				`SELECT status::text FROM pubsub_activities WHERE msg_id=$1 AND service=$2`,
				msgID,
				service,
			).Scan(&status); statusErr != nil {
				return "", false, fmt.Errorf("pubsub_activities inspect claim (%s/%s): %w", service, msgID, statusErr)
			}
			if status == StatusProcessed {
				return "", false, nil
			}
			return "", false, fmt.Errorf("%w: %s/%s", ErrInProgress, service, msgID)
		}
		return "", false, fmt.Errorf("pubsub_activities claim (%s/%s): %w", service, msgID, err)
	}
	return gotToken, true, nil
}

// MarkProcessedLease completes only the caller's current fenced lease.
func MarkProcessedLease(ctx context.Context, repo db.Repository, msgID, service, claimToken string) error {
	const q = `
UPDATE pubsub_activities
   SET status='processed', last_error='', updated_at=now()
 WHERE msg_id=$1 AND service=$2 AND status='processing' AND claim_token=$3::uuid`
	result, err := repo.Exec(ctx, q, msgID, service, claimToken)
	if err != nil {
		return fmt.Errorf("pubsub_activities mark processed (%s/%s): %w", service, msgID, err)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("%w: %s/%s", ErrClaimLost, service, msgID)
	}
	return nil
}

// MarkFailedLease fails only the caller's current fenced lease.
func MarkFailedLease(
	ctx context.Context,
	repo db.Repository,
	msgID, service, claimToken, lastErr string,
) error {
	if len(lastErr) > maxLastError {
		lastErr = lastErr[:maxLastError]
	}
	const q = `
UPDATE pubsub_activities
   SET status='failed', last_error=$4, updated_at=now()
 WHERE msg_id=$1 AND service=$2 AND status='processing' AND claim_token=$3::uuid`
	result, err := repo.Exec(ctx, q, msgID, service, claimToken, lastErr)
	if err != nil {
		return fmt.Errorf("pubsub_activities mark failed (%s/%s): %w", service, msgID, err)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("%w: %s/%s", ErrClaimLost, service, msgID)
	}
	return nil
}

// MarkProcessed records successful handling; future deliveries of this message
// will be skipped by Claim.
func MarkProcessed(ctx context.Context, repo db.Repository, msgID, service string) error {
	const q = `
UPDATE pubsub_activities
   SET status='processed', last_error='', updated_at=now()
 WHERE msg_id=$1 AND service=$2`
	if _, err := repo.Exec(ctx, q, msgID, service); err != nil {
		return fmt.Errorf("pubsub_activities mark processed (%s/%s): %w", service, msgID, err)
	}
	return nil
}

// MarkFailed records a failed attempt while keeping the row as an audit trail.
// The row stays retryable: a subsequent redelivery re-claims it via Claim.
func MarkFailed(ctx context.Context, repo db.Repository, msgID, service, lastErr string) error {
	if len(lastErr) > maxLastError {
		lastErr = lastErr[:maxLastError]
	}
	const q = `
UPDATE pubsub_activities
   SET status='failed', last_error=$3, updated_at=now()
 WHERE msg_id=$1 AND service=$2`
	if _, err := repo.Exec(ctx, q, msgID, service, lastErr); err != nil {
		return fmt.Errorf("pubsub_activities mark failed (%s/%s): %w", service, msgID, err)
	}
	return nil
}
