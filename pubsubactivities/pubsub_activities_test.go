package pubsubactivities

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type rowFunc func(dest ...any) error

func (fn rowFunc) Scan(dest ...any) error { return fn(dest...) }

type claimRepository struct {
	queryRows []pgx.Row
	execTag   pgconn.CommandTag
}

func (repo *claimRepository) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (repo *claimRepository) QueryRow(context.Context, string, ...any) pgx.Row {
	row := repo.queryRows[0]
	repo.queryRows = repo.queryRows[1:]

	return row
}

func (repo *claimRepository) Exec(context.Context, string, ...any) (pgconn.CommandTag, error) {
	return repo.execTag, nil
}

func TestClaimLeaseReturnsErrInProgressForLiveOwner(t *testing.T) {
	t.Parallel()

	repo := &claimRepository{queryRows: []pgx.Row{
		rowFunc(func(...any) error { return pgx.ErrNoRows }),
		rowFunc(func(dest ...any) error {
			status, ok := dest[0].(*string)
			if !ok {
				return errors.New("status destination is not *string")
			}
			*status = StatusProcessing

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
