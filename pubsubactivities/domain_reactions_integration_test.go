package pubsubactivities

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestDomainReactionPostgresIntegration(t *testing.T) {
	databaseURL := os.Getenv("PUBSUBACTIVITIES_POSTGRES_URL")
	if databaseURL == "" {
		t.Skip("PUBSUBACTIVITIES_POSTGRES_URL is not set")
	}

	ctx := context.Background()
	pool := newIsolatedReactionPool(t, ctx, databaseURL)

	t.Run("concurrent claims commit one activity and one delivery", func(t *testing.T) {
		service := "concurrent-domain-events"
		start := make(chan struct{})
		type outcome struct {
			messageID string
			claimed   bool
			err       error
		}
		outcomes := make(chan outcome, 2)
		var wait sync.WaitGroup
		for _, messageID := range []string{"concurrent-message-a", "concurrent-message-b"} {
			messageID := messageID
			wait.Add(1)
			go func() {
				defer wait.Done()
				<-start
				_, claimed, err := ClaimDomainReaction(
					ctx, pool, testEventID, service, testTopic, messageID, 1,
				)
				outcomes <- outcome{messageID: messageID, claimed: claimed, err: err}
			}()
		}
		close(start)
		wait.Wait()
		close(outcomes)

		claimedCount := 0
		inProgressCount := 0
		for result := range outcomes {
			if result.claimed && result.err == nil {
				claimedCount++
				continue
			}
			if !result.claimed && errors.Is(result.err, ErrInProgress) {
				inProgressCount++
				continue
			}
			t.Fatalf("unexpected concurrent result for %s: claimed=%v err=%v", result.messageID, result.claimed, result.err)
		}
		if claimedCount != 1 || inProgressCount != 1 {
			t.Fatalf("claimed=%d in_progress=%d, want 1/1", claimedCount, inProgressCount)
		}
		assertRowCounts(t, ctx, pool, service, 1, 1)
	})

	t.Run("delivery conflict rolls the activity claim back before command", func(t *testing.T) {
		service := "conflict-domain-events"
		messageID := "reused-transport-identity"
		if _, err := pool.Exec(ctx, `
INSERT INTO pubsub_domain_event_deliveries
    (service, event_id, pubsub_message_id, delivery_attempt)
VALUES ($1, $2::uuid, $3, 2)
`, service, testOtherID, messageID); err != nil {
			t.Fatalf("seed conflicting delivery: %v", err)
		}

		var commandCalls atomic.Int32
		_, claimed, err := ClaimDomainReaction(
			ctx, pool, testEventID, service, testTopic, messageID, 4,
		)
		if err == nil && claimed {
			commandCalls.Add(1)
		}
		if !errors.Is(err, ErrDeliveryIdentityConflict) || claimed {
			t.Fatalf("ClaimDomainReaction() claimed=%v err=%v", claimed, err)
		}
		if commandCalls.Load() != 0 {
			t.Fatalf("command calls = %d, want 0", commandCalls.Load())
		}
		assertRowCounts(t, ctx, pool, service, 0, 1)

		var eventID string
		var attempt int
		if err := pool.QueryRow(ctx, `
SELECT event_id::text, delivery_attempt
  FROM pubsub_domain_event_deliveries
 WHERE service=$1 AND pubsub_message_id=$2
`, service, messageID).Scan(&eventID, &attempt); err != nil {
			t.Fatalf("inspect conflicting delivery: %v", err)
		}
		if eventID != testOtherID || attempt != 2 {
			t.Fatalf("conflicting delivery event=%s attempt=%d, want %s/2", eventID, attempt, testOtherID)
		}
	})

	t.Run("live lease cannot resolve and expired lease can", func(t *testing.T) {
		service := "resolution-domain-events"
		_, claimed, err := ClaimDomainReaction(
			ctx, pool, testEventID, service, testTopic, "resolution-message", 1,
		)
		if err != nil || !claimed {
			t.Fatalf("ClaimDomainReaction() claimed=%v err=%v", claimed, err)
		}
		resolved, err := ResolveDomainReaction(ctx, pool, testEventID, service)
		if !errors.Is(err, ErrInProgress) || resolved {
			t.Fatalf("live ResolveDomainReaction() resolved=%v err=%v", resolved, err)
		}
		if _, err := pool.Exec(ctx, `
UPDATE pubsub_activities
   SET updated_at=now() - interval '11 minutes'
 WHERE msg_id=$1 AND service=$2
`, testEventID, service); err != nil {
			t.Fatalf("expire lease: %v", err)
		}
		resolved, err = ResolveDomainReaction(ctx, pool, testEventID, service)
		if err != nil || !resolved {
			t.Fatalf("expired ResolveDomainReaction() resolved=%v err=%v", resolved, err)
		}
		assertActivityStatus(t, ctx, pool, testEventID, service, StatusResolved)
	})

	t.Run("failed retries preserve attempts and republished deliveries", func(t *testing.T) {
		service := "retry-domain-events"
		firstMessage := "retry-message-original"
		token, claimed, err := ClaimDomainReaction(
			ctx, pool, testEventID, service, testTopic, firstMessage, 2,
		)
		if err != nil || !claimed {
			t.Fatalf("first claim claimed=%v err=%v", claimed, err)
		}
		if err := MarkFailedLease(ctx, pool, testEventID, service, token, "transient"); err != nil {
			t.Fatalf("MarkFailedLease(): %v", err)
		}

		token, claimed, err = ClaimDomainReaction(
			ctx, pool, testEventID, service, testTopic, firstMessage, 5,
		)
		if err != nil || !claimed {
			t.Fatalf("same-message reclaim claimed=%v err=%v", claimed, err)
		}
		if err := MarkFailedLease(ctx, pool, testEventID, service, token, "still transient"); err != nil {
			t.Fatalf("second MarkFailedLease(): %v", err)
		}

		token, claimed, err = ClaimDomainReaction(
			ctx, pool, testEventID, service, testTopic, "retry-message-republished", 1,
		)
		if err != nil || !claimed {
			t.Fatalf("republished reclaim claimed=%v err=%v", claimed, err)
		}
		if err := MarkReactedLease(ctx, pool, testEventID, service, token); err != nil {
			t.Fatalf("MarkReactedLease(): %v", err)
		}

		var originalAttempt int
		if err := pool.QueryRow(ctx, `
SELECT delivery_attempt
  FROM pubsub_domain_event_deliveries
 WHERE service=$1 AND pubsub_message_id=$2
`, service, firstMessage).Scan(&originalAttempt); err != nil {
			t.Fatalf("inspect original attempt: %v", err)
		}
		if originalAttempt != 5 {
			t.Fatalf("original attempt = %d, want 5", originalAttempt)
		}
		assertRowCounts(t, ctx, pool, service, 1, 2)
		assertActivityStatus(t, ctx, pool, testEventID, service, StatusReacted)
	})

	t.Run("rejections remain activity-only and monotonic", func(t *testing.T) {
		service := "rejection-domain-events"
		if err := RecordTransportRejection(
			ctx, pool, "malformed-message", service, testTopic, 2, "malformed",
		); err != nil {
			t.Fatalf("RecordTransportRejection(): %v", err)
		}
		if err := RecordTransportRejection(
			ctx, pool, "malformed-message", service, testTopic, 7, "still malformed",
		); err != nil {
			t.Fatalf("repeat RecordTransportRejection(): %v", err)
		}
		if err := RecordApplicationRejection(
			ctx, pool, testOtherID, service, testTopic, 3, "invalid contract",
		); err != nil {
			t.Fatalf("RecordApplicationRejection(): %v", err)
		}
		assertRowCounts(t, ctx, pool, service, 2, 0)

		var attempt int
		var lastError string
		if err := pool.QueryRow(ctx, `
SELECT attempt, last_error
  FROM pubsub_activities
 WHERE msg_id=$1 AND service=$2
`, transportPrefix+"malformed-message", service).Scan(&attempt, &lastError); err != nil {
			t.Fatalf("inspect transport rejection: %v", err)
		}
		if attempt != 7 || lastError != "still malformed" {
			t.Fatalf("transport rejection attempt=%d error=%q", attempt, lastError)
		}
	})
}

func newIsolatedReactionPool(
	t *testing.T,
	ctx context.Context,
	databaseURL string,
) *pgxpool.Pool {
	t.Helper()

	admin, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("connect admin pool: %v", err)
	}
	t.Cleanup(admin.Close)

	schema := "pubsubactivities_" + strings.ReplaceAll(uuid.NewString(), "-", "")
	quotedSchema := pgx.Identifier{schema}.Sanitize()
	if _, err := admin.Exec(ctx, "CREATE SCHEMA "+quotedSchema); err != nil {
		t.Fatalf("create schema: %v", err)
	}
	t.Cleanup(func() {
		_, _ = admin.Exec(context.Background(), "DROP SCHEMA "+quotedSchema+" CASCADE")
	})

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		t.Fatalf("parse database URL: %v", err)
	}
	config.ConnConfig.RuntimeParams["search_path"] = schema
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Fatalf("connect isolated pool: %v", err)
	}
	t.Cleanup(pool.Close)

	if _, err := pool.Exec(ctx, `
CREATE TYPE pubsub_activity_status_t AS ENUM (
    'processing', 'processed', 'failed', 'reacted', 'rejected', 'resolved'
);
CREATE TABLE pubsub_activities (
    msg_id varchar(255) NOT NULL,
    service varchar(64) NOT NULL,
    status pubsub_activity_status_t NOT NULL DEFAULT 'processing',
    last_error text NOT NULL DEFAULT '',
    topic varchar(128) NOT NULL DEFAULT '',
    attempt integer NOT NULL DEFAULT 0,
    claim_token uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (msg_id, service)
);
CREATE TABLE pubsub_domain_event_deliveries (
    service text NOT NULL,
    event_id uuid NOT NULL,
    pubsub_message_id text NOT NULL,
    delivery_attempt integer NOT NULL,
    received_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (service, pubsub_message_id)
);
`); err != nil {
		t.Fatalf("create isolated schema objects: %v", err)
	}
	return pool
}

func assertRowCounts(
	t *testing.T,
	ctx context.Context,
	pool *pgxpool.Pool,
	service string,
	wantActivities, wantDeliveries int,
) {
	t.Helper()
	var activities, deliveries int
	if err := pool.QueryRow(
		ctx,
		"SELECT count(*) FROM pubsub_activities WHERE service=$1",
		service,
	).Scan(&activities); err != nil {
		t.Fatalf("count activities: %v", err)
	}
	if err := pool.QueryRow(
		ctx,
		"SELECT count(*) FROM pubsub_domain_event_deliveries WHERE service=$1",
		service,
	).Scan(&deliveries); err != nil {
		t.Fatalf("count deliveries: %v", err)
	}
	if activities != wantActivities || deliveries != wantDeliveries {
		t.Fatalf(
			"service %s activities/deliveries=%d/%d, want %d/%d",
			service,
			activities,
			deliveries,
			wantActivities,
			wantDeliveries,
		)
	}
}

func assertActivityStatus(
	t *testing.T,
	ctx context.Context,
	pool *pgxpool.Pool,
	eventID, service string,
	want Status,
) {
	t.Helper()
	var got string
	if err := pool.QueryRow(ctx, `
SELECT status::text
  FROM pubsub_activities
 WHERE msg_id=$1 AND service=$2
`, eventID, service).Scan(&got); err != nil {
		t.Fatalf("inspect activity status: %v", err)
	}
	if got != string(want) {
		t.Fatalf("activity status = %q, want %q", got, want)
	}
}

func ExampleClaimDomainReaction() {
	fmt.Println("plan action, claim reaction atomically, then invoke command")
	// Output: plan action, claim reaction atomically, then invoke command
}
