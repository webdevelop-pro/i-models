// Package pubsubactivities is the message-level dedup store for Pub/Sub
// consumers. One row per (msg_id, service) tracks whether a delivery is being
// processed, has been processed, or failed — so a redelivered or
// re-published message is processed at-most-once in effect while failures
// remain retryable and auditable (last_error is kept).
package pubsubactivities

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/webdevelop-pro/go-common/db"
)

// Status values mirror the pubsub_activity_status_t enum.
const (
	StatusProcessing = "processing"
	StatusProcessed  = "processed"
	StatusFailed     = "failed"

	// maxLastError caps stored error text so a pathological error string
	// can't bloat the row.
	maxLastError = 4000
)

// Claim atomically acquires processing ownership of (msgID, service).
//
// Returns true when the caller now owns the message: either it was never seen,
// or a previous attempt left it 'failed' (a legitimate retry re-claims it).
// Returns false when a row already exists as 'processing' (another delivery is
// in-flight) or 'processed' (duplicate of a finished message) — the caller
// should ack and skip. The single statement is race-safe under concurrent
// redeliveries via ON CONFLICT.
func Claim(ctx context.Context, repo db.Repository, msgID, service, topic string, attempt int) (bool, error) {
	const q = `
INSERT INTO pubsub_activities (msg_id, service, status, topic, attempt)
VALUES ($1, $2, 'processing', $3, $4)
ON CONFLICT (msg_id, service) DO UPDATE
   SET status     = 'processing',
       attempt    = EXCLUDED.attempt,
       updated_at = now()
   WHERE pubsub_activities.status = 'failed'
RETURNING msg_id`

	var got string
	err := repo.QueryRow(ctx, q, msgID, service, topic, attempt).Scan(&got)
	if err != nil {
		// No row returned ⇒ the conflicting row was 'processing' or
		// 'processed' (the WHERE excluded it). Not an error: skip.
		if err.Error() == pgx.ErrNoRows.Error() {
			return false, nil
		}
		return false, fmt.Errorf("pubsub_activities claim (%s/%s): %w", service, msgID, err)
	}
	return true, nil
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
