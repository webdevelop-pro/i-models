package pubsubactivities

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/global-torque/go-common/logger/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type failedDeliveryRepo struct {
	query string
	args  []any
	err   error
}

func (r *failedDeliveryRepo) Exec(_ context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	r.query, r.args = query, args
	return pgconn.NewCommandTag("INSERT 0 1"), r.err
}

func (*failedDeliveryRepo) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (*failedDeliveryRepo) QueryRow(context.Context, string, ...any) pgx.Row {
	panic("unexpected QueryRow")
}

func (*failedDeliveryRepo) Begin(context.Context) (pgx.Tx, error) {
	return nil, errors.New("unexpected Begin")
}

func (*failedDeliveryRepo) Lg() logger.Logger { return logger.Logger{} }

func TestRecordFailedDelivery(t *testing.T) {
	t.Parallel()

	repo := &failedDeliveryRepo{}
	longErr := strings.Repeat("x", maxLastError+20)
	err := RecordFailedDelivery(
		context.Background(), repo, "message-1", "wallet-domain-events", "domain-events", 4, longErr,
	)
	if err != nil {
		t.Fatalf("RecordFailedDelivery() error = %v", err)
	}
	if !strings.Contains(repo.query, "status     = 'failed'") {
		t.Fatalf("query does not upsert failed status: %s", repo.query)
	}
	if got := repo.args; len(got) != 5 || got[0] != "message-1" || got[3] != 4 {
		t.Fatalf("unexpected args: %#v", got)
	}
	if got := repo.args[4].(string); len(got) != maxLastError {
		t.Fatalf("stored error length = %d, want %d", len(got), maxLastError)
	}
}

func TestRecordFailedDeliveryWrapsRepositoryError(t *testing.T) {
	t.Parallel()

	repo := &failedDeliveryRepo{err: errors.New("database unavailable")}
	err := RecordFailedDelivery(context.Background(), repo, "message-1", "wallet", "events", 1, "invalid")
	if err == nil || !strings.Contains(err.Error(), "record failed pubsub delivery") {
		t.Fatalf("RecordFailedDelivery() error = %v", err)
	}
}

type rowFunc func(dest ...any) error

func (fn rowFunc) Scan(dest ...any) error { return fn(dest...) }

type claimRepository struct {
	queryRows []pgx.Row
	execTag   pgconn.CommandTag
	queries   []string
	args      [][]any
}

func (repo *claimRepository) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (repo *claimRepository) QueryRow(_ context.Context, query string, args ...any) pgx.Row {
	repo.queries = append(repo.queries, query)
	repo.args = append(repo.args, args)
	row := repo.queryRows[0]
	repo.queryRows = repo.queryRows[1:]

	return row
}

func TestClaimLeaseUsesMonotonicAttemptUpsert(t *testing.T) {
	t.Parallel()

	repo := &claimRepository{queryRows: []pgx.Row{
		rowFunc(func(dest ...any) error {
			claimToken := dest[0].(*string)
			*claimToken = "6d0aaf23-d5ea-4ed5-b020-60fb9ba72155"
			return nil
		}),
	}}

	token, claimed, err := ClaimLease(context.Background(), repo, "event-1", "email", "domain-events", 2)
	if err != nil || !claimed || token == "" {
		t.Fatalf("ClaimLease token=%q claimed=%v error=%v", token, claimed, err)
	}
	if len(repo.queries) != 1 || !strings.Contains(repo.queries[0], "GREATEST(pubsub_activities.attempt, EXCLUDED.attempt)") {
		t.Fatalf("claim query does not preserve the greatest attempt: %v", repo.queries)
	}
	if len(repo.args[0]) < 4 || repo.args[0][3] != 2 {
		t.Fatalf("claim attempt args = %#v, want attempt 2", repo.args[0])
	}
}

func (repo *claimRepository) Exec(context.Context, string, ...any) (pgconn.CommandTag, error) {
	return repo.execTag, nil
}

func (repo *claimRepository) Begin(context.Context) (pgx.Tx, error) {
	return nil, errors.New("unexpected Begin")
}

func (repo *claimRepository) Lg() logger.Logger { return logger.Logger{} }

func TestClaimLeaseReturnsErrInProgressForLiveOwner(t *testing.T) {
	t.Parallel()

	repo := &claimRepository{queryRows: []pgx.Row{
		rowFunc(func(...any) error { return pgx.ErrNoRows }),
		rowFunc(func(dest ...any) error {
			status, ok := dest[0].(*string)
			if !ok {
				return errors.New("status destination is not *string")
			}
			*status = string(StatusProcessing)

			return nil
		}),
	}}

	token, claimed, err := ClaimLease(context.Background(), repo, "event-1", "email", "topic", 2)
	if !errors.Is(err, ErrInProgress) {
		t.Fatalf("ClaimLease error = %v, want ErrInProgress", err)
	}
	if claimed || token != "" {
		t.Fatalf("live concurrent delivery claimed=%v token=%q", claimed, token)
	}
}

func TestTerminalLeaseWriteRejectsStaleToken(t *testing.T) {
	t.Parallel()

	repo := &claimRepository{execTag: pgconn.NewCommandTag("UPDATE 0")}

	err := MarkProcessedLease(context.Background(), repo, "event-1", "email", "stale-token")
	if !errors.Is(err, ErrClaimLost) {
		t.Fatalf("MarkProcessedLease error = %v, want ErrClaimLost", err)
	}

	err = MarkFailedLease(context.Background(), repo, "event-1", "email", "stale-token", "boom")
	if !errors.Is(err, ErrClaimLost) {
		t.Fatalf("MarkFailedLease error = %v, want ErrClaimLost", err)
	}
}
