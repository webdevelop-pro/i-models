package models

import (
	"context"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	commondb "github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
)

type queryTestModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (m queryTestModel) Fields() []string { return DefaultFields(&m) }
func (m queryTestModel) Table() string    { return "test_models" }
func (m *queryTestModel) SetID(id any)    { m.ID = id.(int) }

type stubRepository struct {
	queryRow func(ctx context.Context, sql string, args ...any) pgx.Row
	exec     func(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func (r stubRepository) Query(context.Context, string, ...any) (pgx.Rows, error) { return nil, nil }
func (r stubRepository) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return r.exec(ctx, sql, args...)
}
func (r stubRepository) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return r.queryRow(ctx, sql, args...)
}
func (r stubRepository) Lg() logger.Logger { return logger.Logger{} }

type stubRow struct {
	values []any
	err    error
}

func (r stubRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	for i := range dest {
		switch d := dest[i].(type) {
		case *int:
			*d = r.values[i].(int)
		default:
			panic("unsupported destination type")
		}
	}
	return nil
}

var _ commondb.Repository = stubRepository{}

func TestExistsAcceptsMapAndExpressions(t *testing.T) {
	var gotSQL string
	var gotArgs []any
	repo := stubRepository{
		queryRow: func(_ context.Context, sql string, args ...any) pgx.Row {
			gotSQL = sql
			gotArgs = args
			return stubRow{values: []any{1}}
		},
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}

	exists, err := Exists[queryTestModel](context.Background(), repo,
		map[string]any{"name": "alice"},
		sq.Expr("lower(email) = ?", "alice@example.com"),
		sq.NotEq{"id": 42},
	)
	if err != nil {
		t.Fatalf("Exists returned error: %v", err)
	}
	if !exists {
		t.Fatalf("Exists returned false, want true")
	}

	wantSQL := "SELECT 1 FROM test_models WHERE name = $1 AND lower(email) = $2 AND id <> $3"
	if gotSQL != wantSQL {
		t.Fatalf("unexpected SQL:\n got: %s\nwant: %s", gotSQL, wantSQL)
	}
	if len(gotArgs) != 3 || gotArgs[0] != "alice" || gotArgs[1] != "alice@example.com" || gotArgs[2] != 42 {
		t.Fatalf("unexpected args: %#v", gotArgs)
	}
}

func TestDeleteAcceptsMapAndExpressions(t *testing.T) {
	var gotSQL string
	var gotArgs []any
	repo := stubRepository{
		queryRow: func(_ context.Context, _ string, _ ...any) pgx.Row {
			return stubRow{err: pgx.ErrNoRows}
		},
		exec: func(_ context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
			gotSQL = sql
			gotArgs = args
			return pgconn.NewCommandTag("DELETE 1"), nil
		},
	}

	deleted, err := Delete[queryTestModel](context.Background(), repo,
		map[string]any{"name": "alice"},
		sq.Expr("archived_at IS NULL"),
		sq.NotEq{"id": 42},
	)
	if err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}
	if !deleted {
		t.Fatalf("Delete returned false, want true")
	}

	wantSQL := "DELETE FROM test_models WHERE name = $1 AND archived_at IS NULL AND id <> $2"
	if gotSQL != wantSQL {
		t.Fatalf("unexpected SQL:\n got: %s\nwant: %s", gotSQL, wantSQL)
	}
	if len(gotArgs) != 2 || gotArgs[0] != "alice" || gotArgs[1] != 42 {
		t.Fatalf("unexpected args: %#v", gotArgs)
	}
}
