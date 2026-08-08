// Package pubsubactivities is the message-level dedup and domain-reaction
// audit store for Pub/Sub consumers.
package pubsubactivities

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/global-torque/go-common/db/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	// ErrInProgress tells push handlers to NACK a concurrent delivery. ACKing it
	// would acknowledge the original delivery too, creating a crash-loss window.
	ErrInProgress = errors.New("pubsub activity is already processing")

	// ErrClaimLost means a stale worker tried to complete a lease that has since
	// been reclaimed. The caller must not overwrite the current owner's state.
	ErrClaimLost = errors.New("pubsub activity processing lease was lost")

	// ErrUnknownStatus means persisted activity state is newer than or corrupt
	// for this library. Failing closed prevents it from being reclaimed as work.
	ErrUnknownStatus = errors.New("unknown pubsub activity status")

	// ErrInvalidIdentity means an audit identity is empty, malformed, or wider
	// than the canonical pubsub_activities schema accepts.
	ErrInvalidIdentity = errors.New("invalid pubsub activity identity")

	// ErrDeliveryIdentityConflict means a Pub/Sub message identity is already
	// attached to another immutable outbox event for the same service.
	ErrDeliveryIdentityConflict = errors.New("pubsub delivery identity belongs to another event")
)

// Status mirrors pubsub_activity_status_t.
type Status string

const (
	StatusProcessing Status = "processing"
	StatusProcessed  Status = "processed"
	StatusFailed     Status = "failed"
	StatusReacted    Status = "reacted"
	StatusRejected   Status = "rejected"
	StatusResolved   Status = "resolved"

	maxMessageIDLength = 255
	maxServiceLength   = 64
	maxTopicLength     = 128
	maxLastError       = 4000
	processingLease    = "10 minutes"
	terminalTimeout    = 5 * time.Second
	transportPrefix    = "transport:"
)

// IsTerminal reports whether a status permanently prevents another claim.
func (status Status) IsTerminal() bool {
	switch status {
	case StatusProcessed, StatusReacted, StatusRejected, StatusResolved:
		return true
	default:
		return false
	}
}

// IsValid reports whether a status is part of the shared activity contract.
func (status Status) IsValid() bool {
	switch status {
	case StatusProcessing, StatusProcessed, StatusFailed,
		StatusReacted, StatusRejected, StatusResolved:
		return true
	default:
		return false
	}
}

// TransactionalRepository is the narrow query surface required to atomically
// claim a domain reaction and record its transport delivery.
type TransactionalRepository interface {
	db.Repository
	Begin(context.Context) (pgx.Tx, error)
}

// RecordDomainDelivery stores transport identity independently from an outbox
// event id.
//
// Deprecated: domain-event consumers must use ClaimDomainReaction so the
// activity claim and delivery row commit atomically after application planning.
// This compatibility function remains until every transport-boundary caller is
// migrated.
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

// RecordFailedDelivery retains the legacy generic failed-delivery behavior.
// Compatibility-only service builds rely on this function not writing new enum
// values before the additive database migration has been applied.
//
// Deprecated: migrated domain transports must use RecordTransportRejection;
// parsed application-contract failures must use RecordApplicationRejection.
func RecordFailedDelivery(
	ctx context.Context,
	repo db.Repository,
	msgID, service, topic string,
	attempt int,
	lastErr string,
) error {
	if msgID == "" {
		return nil
	}
	lastErr = truncateError(lastErr)

	_, err := repo.Exec(ctx, `
INSERT INTO pubsub_activities (msg_id, service, status, topic, attempt, last_error)
VALUES ($1, $2, 'failed', $3, $4, $5)
ON CONFLICT (msg_id, service) DO UPDATE
   SET status     = 'failed',
       topic      = EXCLUDED.topic,
       attempt    = GREATEST(pubsub_activities.attempt, EXCLUDED.attempt),
       last_error = EXCLUDED.last_error,
       updated_at = now()
`, msgID, service, topic, attempt, lastErr)
	if err != nil {
		return fmt.Errorf("record failed pubsub delivery (%s/%s): %w", service, msgID, err)
	}
	return nil
}

// Claim atomically acquires processing ownership of (msgID, service).
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
	return claimLease(ctx, repo, msgID, service, topic, attempt)
}

func claimLease(
	ctx context.Context,
	repo db.Repository,
	msgID, service, topic string,
	attempt int,
) (string, bool, error) {
	const q = `
INSERT INTO pubsub_activities (msg_id, service, status, topic, attempt, claim_token)
VALUES ($1, $2, 'processing', $3, $4, $6::uuid)
ON CONFLICT (msg_id, service) DO UPDATE
   SET status      = 'processing',
       topic       = EXCLUDED.topic,
       attempt     = GREATEST(pubsub_activities.attempt, EXCLUDED.attempt),
       last_error  = '',
       claim_token = EXCLUDED.claim_token,
       updated_at  = now()
   WHERE pubsub_activities.status = 'failed'
      OR (pubsub_activities.status = 'processing'
          AND pubsub_activities.updated_at < now() - $5::interval)
RETURNING claim_token::text`

	claimToken := uuid.NewString()
	var gotToken string
	err := repo.QueryRow(ctx, q, msgID, service, topic, attempt, processingLease, claimToken).Scan(&gotToken)
	if err == nil {
		return gotToken, true, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return "", false, fmt.Errorf("pubsub_activities claim (%s/%s): %w", service, msgID, err)
	}

	var rawStatus string
	if statusErr := repo.QueryRow(
		ctx,
		`SELECT status::text FROM pubsub_activities WHERE msg_id=$1 AND service=$2`,
		msgID,
		service,
	).Scan(&rawStatus); statusErr != nil {
		return "", false, fmt.Errorf("pubsub_activities inspect claim (%s/%s): %w", service, msgID, statusErr)
	}

	status := Status(rawStatus)
	if !status.IsValid() {
		return "", false, fmt.Errorf("%w %q for %s/%s", ErrUnknownStatus, rawStatus, service, msgID)
	}
	if status.IsTerminal() {
		return "", false, nil
	}
	if status == StatusProcessing {
		return "", false, fmt.Errorf("%w: %s/%s", ErrInProgress, service, msgID)
	}

	// A failed row is eligible in the upsert above. Reaching inspection with it
	// means a database rule prevented acquisition, so fail closed rather than
	// pretending another worker owns a live lease.
	return "", false, fmt.Errorf("pubsub_activities failed claim remained claimable: %s/%s", service, msgID)
}

// ClaimDomainReaction atomically acquires a fenced event-id claim and records
// the Pub/Sub delivery before any business command may begin.
func ClaimDomainReaction(
	ctx context.Context,
	repo TransactionalRepository,
	eventID, service, topic, pubsubMessageID string,
	attempt int,
) (string, bool, error) {
	if err := validateDomainIdentity(eventID, service, topic, pubsubMessageID); err != nil {
		return "", false, err
	}

	tx, err := repo.Begin(ctx)
	if err != nil {
		return "", false, fmt.Errorf("begin domain reaction claim (%s/%s): %w", service, eventID, err)
	}
	defer rollback(tx)

	claimToken, claimed, err := claimLease(ctx, tx, eventID, service, topic, attempt)
	if err != nil || !claimed {
		return "", claimed, err
	}

	const deliveryQuery = `
INSERT INTO pubsub_domain_event_deliveries
    (service, event_id, pubsub_message_id, delivery_attempt)
VALUES ($1, $2::uuid, $3, $4)
ON CONFLICT (service, pubsub_message_id) DO UPDATE
   SET delivery_attempt = GREATEST(
           pubsub_domain_event_deliveries.delivery_attempt,
           EXCLUDED.delivery_attempt
       )
 WHERE pubsub_domain_event_deliveries.event_id = EXCLUDED.event_id
RETURNING event_id::text`

	var recordedEventID string
	err = tx.QueryRow(ctx, deliveryQuery, service, eventID, pubsubMessageID, attempt).Scan(&recordedEventID)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, fmt.Errorf("%w: %s/%s", ErrDeliveryIdentityConflict, service, pubsubMessageID)
	}
	if err != nil {
		return "", false, fmt.Errorf("record claimed domain delivery (%s/%s): %w", service, pubsubMessageID, err)
	}
	if recordedEventID != eventID {
		return "", false, fmt.Errorf("%w: %s/%s", ErrDeliveryIdentityConflict, service, pubsubMessageID)
	}
	if err := tx.Commit(ctx); err != nil {
		return "", false, fmt.Errorf("commit domain reaction claim (%s/%s): %w", service, eventID, err)
	}
	return claimToken, true, nil
}

// ResolveDomainReaction closes an existing failed or expired claimed reaction
// that authoritative planning proves no longer requires a command. It never
// creates activity or delivery state for a first-time no-action outcome.
func ResolveDomainReaction(
	ctx context.Context,
	repo TransactionalRepository,
	eventID, service string,
) (bool, error) {
	if err := validateRequiredIdentity("event id", eventID, maxMessageIDLength); err != nil {
		return false, err
	}
	if _, err := uuid.Parse(eventID); err != nil {
		return false, fmt.Errorf("%w: event id is not a UUID", ErrInvalidIdentity)
	}
	if err := validateRequiredIdentity("service", service, maxServiceLength); err != nil {
		return false, err
	}

	tx, err := repo.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("begin domain reaction resolution (%s/%s): %w", service, eventID, err)
	}
	defer rollback(tx)

	const resolutionQuery = `
UPDATE pubsub_activities AS activity
   SET status='resolved', last_error='', updated_at=now()
 WHERE activity.msg_id=$1::text
   AND activity.service=$2
   AND (activity.status='failed'
        OR (activity.status='processing'
            AND activity.updated_at < now() - $3::interval))
   AND EXISTS (
       SELECT 1
         FROM pubsub_domain_event_deliveries AS delivery
        WHERE delivery.event_id=$1::uuid
          AND delivery.service=$2
   )
RETURNING activity.status::text`

	var status string
	err = tx.QueryRow(ctx, resolutionQuery, eventID, service, processingLease).Scan(&status)
	if err == nil {
		if err := tx.Commit(ctx); err != nil {
			return false, fmt.Errorf("commit domain reaction resolution (%s/%s): %w", service, eventID, err)
		}
		return true, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("resolve domain reaction (%s/%s): %w", service, eventID, err)
	}

	var rawStatus string
	var hasDelivery bool
	err = tx.QueryRow(ctx, `
SELECT activity.status::text,
       EXISTS (
           SELECT 1
             FROM pubsub_domain_event_deliveries AS delivery
            WHERE delivery.event_id=$1::uuid
              AND delivery.service=$2
       )
  FROM pubsub_activities AS activity
 WHERE activity.msg_id=$1::text AND activity.service=$2
`, eventID, service).Scan(&rawStatus, &hasDelivery)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("inspect domain reaction resolution (%s/%s): %w", service, eventID, err)
	}
	if !hasDelivery {
		return false, nil
	}

	current := Status(rawStatus)
	if !current.IsValid() {
		return false, fmt.Errorf("%w %q for %s/%s", ErrUnknownStatus, rawStatus, service, eventID)
	}
	if current == StatusProcessing {
		return false, fmt.Errorf("%w: %s/%s", ErrInProgress, service, eventID)
	}
	return false, nil
}

// RecordTransportRejection stores a permanently invalid trusted transport
// delivery under a namespace that cannot collide with an immutable event ID.
func RecordTransportRejection(
	ctx context.Context,
	repo db.Repository,
	pubsubMessageID, service, topic string,
	attempt int,
	lastErr string,
) error {
	if err := validateRequiredIdentity(
		"Pub/Sub message id",
		pubsubMessageID,
		maxMessageIDLength-len(transportPrefix),
	); err != nil {
		return err
	}
	return recordRejection(
		ctx,
		repo,
		transportPrefix+pubsubMessageID,
		service,
		topic,
		attempt,
		lastErr,
	)
}

// RecordApplicationRejection stores a parsed permanent application-contract
// failure keyed by immutable event ID without creating a delivery row.
func RecordApplicationRejection(
	ctx context.Context,
	repo db.Repository,
	eventID, service, topic string,
	attempt int,
	lastErr string,
) error {
	if err := validateRequiredIdentity("event id", eventID, maxMessageIDLength); err != nil {
		return err
	}
	if _, err := uuid.Parse(eventID); err != nil {
		return fmt.Errorf("%w: event id is not a UUID", ErrInvalidIdentity)
	}
	return recordRejection(ctx, repo, eventID, service, topic, attempt, lastErr)
}

func recordRejection(
	ctx context.Context,
	repo db.Repository,
	msgID, service, topic string,
	attempt int,
	lastErr string,
) error {
	if err := validateRequiredIdentity("service", service, maxServiceLength); err != nil {
		return err
	}
	if err := validateRequiredIdentity("topic", topic, maxTopicLength); err != nil {
		return err
	}
	lastErr = truncateError(lastErr)

	const q = `
INSERT INTO pubsub_activities (msg_id, service, status, topic, attempt, last_error)
VALUES ($1, $2, 'rejected', $3, $4, $5)
ON CONFLICT (msg_id, service) DO UPDATE
   SET topic      = EXCLUDED.topic,
       attempt    = GREATEST(pubsub_activities.attempt, EXCLUDED.attempt),
       last_error = EXCLUDED.last_error,
       updated_at = now()
 WHERE pubsub_activities.status = 'rejected'
RETURNING status::text`

	var status string
	if err := repo.QueryRow(ctx, q, msgID, service, topic, attempt, lastErr).Scan(&status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("record rejection conflicts with existing activity (%s/%s)", service, msgID)
		}
		return fmt.Errorf("record rejected pubsub delivery (%s/%s): %w", service, msgID, err)
	}
	return nil
}

// MarkProcessedLease completes only the caller's current fenced generic lease.
func MarkProcessedLease(ctx context.Context, repo db.Repository, msgID, service, claimToken string) error {
	return markLease(ctx, repo, msgID, service, claimToken, StatusProcessed, "")
}

// MarkReactedLease completes only the current fenced domain-reaction lease.
// Its persistence context is bounded and detached from request cancellation;
// this protects terminal evidence without extending business work.
func MarkReactedLease(ctx context.Context, repo db.Repository, eventID, service, claimToken string) error {
	terminalCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), terminalTimeout)
	defer cancel()
	return markLease(terminalCtx, repo, eventID, service, claimToken, StatusReacted, "")
}

// MarkRejectedLease permanently rejects only the current fenced domain lease.
func MarkRejectedLease(
	ctx context.Context,
	repo db.Repository,
	eventID, service, claimToken, lastErr string,
) error {
	terminalCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), terminalTimeout)
	defer cancel()
	return markLease(terminalCtx, repo, eventID, service, claimToken, StatusRejected, truncateError(lastErr))
}

func markLease(
	ctx context.Context,
	repo db.Repository,
	msgID, service, claimToken string,
	status Status,
	lastErr string,
) error {
	const q = `
UPDATE pubsub_activities
   SET status=$4::pubsub_activity_status_t, last_error=$5, updated_at=now()
 WHERE msg_id=$1 AND service=$2 AND status='processing' AND claim_token=$3::uuid`
	result, err := repo.Exec(ctx, q, msgID, service, claimToken, status, lastErr)
	if err != nil {
		return fmt.Errorf("pubsub_activities mark %s (%s/%s): %w", status, service, msgID, err)
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
	return markLease(ctx, repo, msgID, service, claimToken, StatusFailed, truncateError(lastErr))
}

// MarkProcessed records successful generic handling; future deliveries of this
// message will be skipped by Claim.
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

// MarkFailed records a failed generic attempt while keeping it retryable.
func MarkFailed(ctx context.Context, repo db.Repository, msgID, service, lastErr string) error {
	const q = `
UPDATE pubsub_activities
   SET status='failed', last_error=$3, updated_at=now()
 WHERE msg_id=$1 AND service=$2`
	if _, err := repo.Exec(ctx, q, msgID, service, truncateError(lastErr)); err != nil {
		return fmt.Errorf("pubsub_activities mark failed (%s/%s): %w", service, msgID, err)
	}
	return nil
}

func validateDomainIdentity(eventID, service, topic, pubsubMessageID string) error {
	if err := validateRequiredIdentity("event id", eventID, maxMessageIDLength); err != nil {
		return err
	}
	if _, err := uuid.Parse(eventID); err != nil {
		return fmt.Errorf("%w: event id is not a UUID", ErrInvalidIdentity)
	}
	if err := validateRequiredIdentity("service", service, maxServiceLength); err != nil {
		return err
	}
	if err := validateRequiredIdentity("topic", topic, maxTopicLength); err != nil {
		return err
	}
	return validateRequiredIdentity("Pub/Sub message id", pubsubMessageID, maxMessageIDLength)
}

func validateRequiredIdentity(name, value string, maxLength int) error {
	if value == "" {
		return fmt.Errorf("%w: %s is empty", ErrInvalidIdentity, name)
	}
	if len(value) > maxLength {
		return fmt.Errorf("%w: %s exceeds %d bytes", ErrInvalidIdentity, name, maxLength)
	}
	return nil
}

func truncateError(lastErr string) string {
	lastErr = strings.ToValidUTF8(lastErr, "�")
	if len(lastErr) > maxLastError {
		lastErr = lastErr[:maxLastError]
		for !utf8.ValidString(lastErr) {
			lastErr = lastErr[:len(lastErr)-1]
		}
	}
	return lastErr
}

func rollback(tx pgx.Tx) {
	ctx, cancel := context.WithTimeout(context.Background(), terminalTimeout)
	defer cancel()
	_ = tx.Rollback(ctx)
}
