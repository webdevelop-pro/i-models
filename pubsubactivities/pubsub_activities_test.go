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
