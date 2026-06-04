package models

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	commondb "github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/go-common/queue/pclient"
)

type queryTestModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (m queryTestModel) Fields() []string { return DefaultFields(&m) }
func (m queryTestModel) Table() string    { return "test_models" }
func (m *queryTestModel) SetID(id any)    { m.ID = id.(int) }
func (m *queryTestModel) SetDB(commondb.Repository) {
}

type stubRepository struct {
	query    func(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	queryRow func(ctx context.Context, sql string, args ...any) pgx.Row
	exec     func(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func (r stubRepository) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if r.query == nil {
		return nil, nil
	}
	return r.query(ctx, sql, args...)
}
func (r stubRepository) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return r.exec(ctx, sql, args...)
}
func (r stubRepository) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return r.queryRow(ctx, sql, args...)
}
func (r stubRepository) Lg() logger.Logger { return logger.Logger{} }

type stubRows struct {
	fields  []string
	values  [][]any
	idx     int
	err     error
	scanErr error
}

func (r *stubRows) Close()                        {}
func (r *stubRows) Err() error                    { return r.err }
func (r *stubRows) CommandTag() pgconn.CommandTag { return pgconn.CommandTag{} }
func (r *stubRows) FieldDescriptions() []pgconn.FieldDescription {
	fields := make([]pgconn.FieldDescription, len(r.fields))
	for i, field := range r.fields {
		fields[i] = pgconn.FieldDescription{Name: field}
	}
	return fields
}
func (r *stubRows) Next() bool {
	if r.err != nil || r.idx >= len(r.values) {
		return false
	}
	r.idx++
	return true
}
func (r *stubRows) Scan(dest ...any) error {
	if r.scanErr != nil {
		return r.scanErr
	}
	row := r.values[r.idx-1]
	for i := range dest {
		if row[i] == nil {
			continue
		}
		target := reflect.ValueOf(dest[i])
		if target.Kind() != reflect.Pointer || target.IsNil() {
			continue
		}
		value := reflect.ValueOf(row[i])
		if value.Type().AssignableTo(target.Elem().Type()) {
			target.Elem().Set(value)
		} else if value.Type().ConvertibleTo(target.Elem().Type()) {
			target.Elem().Set(value.Convert(target.Elem().Type()))
		}
	}
	return nil
}
func (r *stubRows) Values() ([]any, error) { return r.values[r.idx-1], nil }
func (r *stubRows) RawValues() [][]byte    { return nil }
func (r *stubRows) Conn() *pgx.Conn        { return nil }

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

type badSqlizer struct {
	err error
}

func (s badSqlizer) ToSql() (string, []interface{}, error) {
	return "", nil, s.err
}

var _ commondb.Repository = stubRepository{}

func TestRetrieveOneBuildsSQLWithSuffix(t *testing.T) {
	var gotSQL string
	var gotArgs []any
	repo := stubRepository{
		query: func(_ context.Context, sql string, args ...any) (pgx.Rows, error) {
			gotSQL = sql
			gotArgs = args
			return &stubRows{
				fields: []string{"id", "name"},
				values: [][]any{
					{7, "alice"},
				},
			}, nil
		},
		queryRow: func(_ context.Context, _ string, _ ...any) pgx.Row {
			return stubRow{err: pgx.ErrNoRows}
		},
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}

	got, err := RetrieveOne[queryTestModel](
		context.Background(),
		repo,
		sq.Eq{"name": "alice"},
		sq.Expr("ORDER BY id LIMIT ?", 1),
	)
	if err != nil {
		t.Fatalf("RetrieveOne returned error: %v", err)
	}

	wantSQL := "SELECT id,name FROM test_models WHERE name = $1 ORDER BY id LIMIT $2"
	if gotSQL != wantSQL {
		t.Fatalf("unexpected SQL:\n got: %s\nwant: %s", gotSQL, wantSQL)
	}
	if len(gotArgs) != 2 || gotArgs[0] != "alice" || gotArgs[1] != 1 {
		t.Fatalf("unexpected args: %#v", gotArgs)
	}
	if got.ID != 7 || got.Name != "alice" {
		t.Fatalf("unexpected model: %#v", got)
	}
}

func TestRetrieveAllBuildsSQLWithSuffix(t *testing.T) {
	var gotSQL string
	var gotArgs []any
	repo := stubRepository{
		query: func(_ context.Context, sql string, args ...any) (pgx.Rows, error) {
			gotSQL = sql
			gotArgs = args
			return &stubRows{
				fields: []string{"id", "name"},
				values: [][]any{
					{7, "alice"},
					{8, "bob"},
				},
			}, nil
		},
		queryRow: func(_ context.Context, _ string, _ ...any) pgx.Row {
			return stubRow{err: pgx.ErrNoRows}
		},
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}

	got, err := RetrieveAll[queryTestModel](
		context.Background(),
		repo,
		sq.Gt{"id": 6},
		sq.Expr("ORDER BY name LIMIT ?", 2),
	)
	if err != nil {
		t.Fatalf("RetrieveAll returned error: %v", err)
	}

	wantSQL := "SELECT id,name FROM test_models WHERE id > $1 ORDER BY name LIMIT $2"
	if gotSQL != wantSQL {
		t.Fatalf("unexpected SQL:\n got: %s\nwant: %s", gotSQL, wantSQL)
	}
	if len(gotArgs) != 2 || gotArgs[0] != 6 || gotArgs[1] != 2 {
		t.Fatalf("unexpected args: %#v", gotArgs)
	}
	if len(got) != 2 || got[0].Name != "alice" || got[1].Name != "bob" {
		t.Fatalf("unexpected models: %#v", got)
	}
}

func TestRetrieveOneQueryError(t *testing.T) {
	queryErr := errors.New("query failed")
	repo := stubRepository{
		query: func(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
			return nil, queryErr
		},
	}

	_, err := RetrieveOne[queryTestModel](context.Background(), repo, sq.Eq{"id": 1})
	if err == nil {
		t.Fatalf("RetrieveOne returned nil error")
	}
	if !errors.Is(err, queryErr) {
		t.Fatalf("RetrieveOne error does not wrap query error: %v", err)
	}
}

func TestRetrieveAllQueryError(t *testing.T) {
	queryErr := errors.New("query failed")
	repo := stubRepository{
		query: func(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
			return nil, queryErr
		},
	}

	_, err := RetrieveAll[queryTestModel](context.Background(), repo, sq.Eq{"id": 1})
	if err == nil {
		t.Fatalf("RetrieveAll returned nil error")
	}
	if !errors.Is(err, queryErr) {
		t.Fatalf("RetrieveAll error does not wrap query error: %v", err)
	}
}

func TestRetrieveOneNotFoundWrapsSentinels(t *testing.T) {
	repo := stubRepository{
		query: func(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
			return &stubRows{fields: []string{"id", "name"}}, nil
		},
	}

	_, err := RetrieveOne[queryTestModel](context.Background(), repo, sq.Eq{"id": 1})
	if err == nil {
		t.Fatalf("RetrieveOne returned nil error")
	}
	if !errors.Is(err, ErrRecordNotFound) {
		t.Fatalf("RetrieveOne error does not wrap ErrRecordNotFound: %v", err)
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("RetrieveOne error does not wrap pgx.ErrNoRows: %v", err)
	}
}

func TestRetrieveOneScanErrorWrapsCause(t *testing.T) {
	scanErr := errors.New("scan failed")
	repo := stubRepository{
		query: func(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
			return &stubRows{
				fields:  []string{"id", "name"},
				values:  [][]any{{7, "alice"}},
				scanErr: scanErr,
			}, nil
		},
	}

	_, err := RetrieveOne[queryTestModel](context.Background(), repo, sq.Eq{"id": 1})
	if err == nil {
		t.Fatalf("RetrieveOne returned nil error")
	}
	if !errors.Is(err, scanErr) {
		t.Fatalf("RetrieveOne error does not wrap scan error: %v", err)
	}
}

func TestRetrieveOnePrepareErrorWrapsCause(t *testing.T) {
	prepareErr := errors.New("prepare failed")
	repo := stubRepository{}

	_, err := RetrieveOne[queryTestModel](context.Background(), repo, badSqlizer{err: prepareErr})
	if err == nil {
		t.Fatalf("RetrieveOne returned nil error")
	}
	if !errors.Is(err, prepareErr) {
		t.Fatalf("RetrieveOne error does not wrap prepare error: %v", err)
	}
}

func TestRetrieveAllScanErrorWrapsCause(t *testing.T) {
	scanErr := errors.New("scan failed")
	repo := stubRepository{
		query: func(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
			return &stubRows{
				fields:  []string{"id", "name"},
				values:  [][]any{{7, "alice"}},
				scanErr: scanErr,
			}, nil
		},
	}

	_, err := RetrieveAll[queryTestModel](context.Background(), repo, sq.Eq{"id": 1})
	if err == nil {
		t.Fatalf("RetrieveAll returned nil error")
	}
	if !errors.Is(err, scanErr) {
		t.Fatalf("RetrieveAll error does not wrap scan error: %v", err)
	}
}

func TestRetrieveAllPrepareErrorWrapsCause(t *testing.T) {
	prepareErr := errors.New("prepare failed")
	repo := stubRepository{}

	_, err := RetrieveAll[queryTestModel](context.Background(), repo, badSqlizer{err: prepareErr})
	if err == nil {
		t.Fatalf("RetrieveAll returned nil error")
	}
	if !errors.Is(err, prepareErr) {
		t.Fatalf("RetrieveAll error does not wrap prepare error: %v", err)
	}
}

func TestCreateQueryRowErrorWrapsCause(t *testing.T) {
	createErr := errors.New("insert failed")
	repo := stubRepository{
		queryRow: func(_ context.Context, _ string, _ ...any) pgx.Row {
			return stubRow{err: createErr}
		},
	}

	_, err := Create[queryTestModel](context.Background(), repo, map[string]any{"name": "alice"})
	if err == nil {
		t.Fatalf("Create returned nil error")
	}
	if !errors.Is(err, createErr) {
		t.Fatalf("Create error does not wrap query row error: %v", err)
	}
}

func TestCreatePrepareErrorWrapsCause(t *testing.T) {
	prepareErr := errors.New("prepare failed")
	repo := stubRepository{}

	_, err := Create[queryTestModel](context.Background(), repo, map[string]any{"name": badSqlizer{err: prepareErr}})
	if err == nil {
		t.Fatalf("Create returned nil error")
	}
	if !errors.Is(err, prepareErr) {
		t.Fatalf("Create error does not wrap prepare error: %v", err)
	}
}

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

func TestExistsQueryRowErrorWrapsCause(t *testing.T) {
	existsErr := errors.New("exists failed")
	repo := stubRepository{
		queryRow: func(_ context.Context, _ string, _ ...any) pgx.Row {
			return stubRow{err: existsErr}
		},
	}

	_, err := Exists[queryTestModel](context.Background(), repo, map[string]any{"name": "alice"})
	if err == nil {
		t.Fatalf("Exists returned nil error")
	}
	if !errors.Is(err, existsErr) {
		t.Fatalf("Exists error does not wrap query row error: %v", err)
	}
}

func TestExistsPrepareErrorWrapsCause(t *testing.T) {
	prepareErr := errors.New("prepare failed")
	repo := stubRepository{}

	_, err := Exists[queryTestModel](context.Background(), repo, map[string]any{"name": "alice"}, badSqlizer{err: prepareErr})
	if err == nil {
		t.Fatalf("Exists returned nil error")
	}
	if !errors.Is(err, prepareErr) {
		t.Fatalf("Exists error does not wrap prepare error: %v", err)
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

func TestDeletePrepareErrorWrapsCause(t *testing.T) {
	prepareErr := errors.New("prepare failed")
	repo := stubRepository{}

	_, err := Delete[queryTestModel](context.Background(), repo, map[string]any{"name": "alice"}, badSqlizer{err: prepareErr})
	if err == nil {
		t.Fatalf("Delete returned nil error")
	}
	if !errors.Is(err, prepareErr) {
		t.Fatalf("Delete error does not wrap prepare error: %v", err)
	}
}

func TestDeleteExecErrorWrapsCause(t *testing.T) {
	deleteErr := errors.New("delete failed")
	repo := stubRepository{
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, deleteErr
		},
	}

	_, err := Delete[queryTestModel](context.Background(), repo, map[string]any{"name": "alice"})
	if err == nil {
		t.Fatalf("Delete returned nil error")
	}
	if !errors.Is(err, deleteErr) {
		t.Fatalf("Delete error does not wrap exec error: %v", err)
	}
}

func TestDeleteRejectsTautologySqlizerPredicates(t *testing.T) {
	tests := []struct {
		name string
		expr sq.Sqlizer
	}{
		{name: "spaced", expr: sq.Expr("1 = 1")},
		{name: "parenthesized", expr: sq.Expr("(1 = 1)")},
		{name: "composed", expr: sq.And{sq.Expr("1 = 1")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := stubRepository{
				exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
					t.Fatalf("Exec should not be called")
					return pgconn.CommandTag{}, nil
				},
			}

			deleted, err := Delete[queryTestModel](context.Background(), repo, nil, tt.expr)
			if err == nil {
				t.Fatalf("Delete returned nil error")
			}
			if deleted {
				t.Fatalf("Delete returned true, want false")
			}
			if !strings.Contains(err.Error(), ErrEmptyPredicate) {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestDeleteRejectsEmptyPredicate(t *testing.T) {
	repo := stubRepository{
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			t.Fatalf("Exec should not be called")
			return pgconn.CommandTag{}, nil
		},
	}

	deleted, err := Delete[queryTestModel](context.Background(), repo, nil)
	if err == nil {
		t.Fatalf("Delete returned nil error")
	}
	if deleted {
		t.Fatalf("Delete returned true, want false")
	}
	if !strings.Contains(err.Error(), ErrEmptyPredicate) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteRejectsEmptySqlizerPredicate(t *testing.T) {
	repo := stubRepository{
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			t.Fatalf("Exec should not be called")
			return pgconn.CommandTag{}, nil
		},
	}

	deleted, err := Delete[queryTestModel](context.Background(), repo, nil, sq.Eq{})
	if err == nil {
		t.Fatalf("Delete returned nil error")
	}
	if deleted {
		t.Fatalf("Delete returned true, want false")
	}
	if !strings.Contains(err.Error(), ErrEmptyPredicate) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateAcceptsMapAndExpressions(t *testing.T) {
	var gotSQL string
	var gotArgs []any
	repo := stubRepository{
		queryRow: func(_ context.Context, _ string, _ ...any) pgx.Row {
			return stubRow{err: pgx.ErrNoRows}
		},
		exec: func(_ context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
			gotSQL = sql
			gotArgs = args
			return pgconn.NewCommandTag("UPDATE 1"), nil
		},
	}

	updated, err := Update[queryTestModel](context.Background(), repo,
		map[string]any{"name": "alice"},
		map[string]any{"name": "bob"},
		sq.Expr("archived_at IS NULL"),
		sq.NotEq{"id": 42},
	)
	if err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	if !updated {
		t.Fatalf("Update returned false, want true")
	}

	wantSQL := "UPDATE test_models SET name = $1 WHERE name = $2 AND archived_at IS NULL AND id <> $3"
	if gotSQL != wantSQL {
		t.Fatalf("unexpected SQL:\n got: %s\nwant: %s", gotSQL, wantSQL)
	}
	if len(gotArgs) != 3 || gotArgs[0] != "bob" || gotArgs[1] != "alice" || gotArgs[2] != 42 {
		t.Fatalf("unexpected args: %#v", gotArgs)
	}
}

func TestUpdatePrepareErrorWrapsCause(t *testing.T) {
	prepareErr := errors.New("prepare failed")
	repo := stubRepository{}

	_, err := Update[queryTestModel](
		context.Background(),
		repo,
		map[string]any{"name": "alice"},
		map[string]any{"name": "bob"},
		badSqlizer{err: prepareErr},
	)
	if err == nil {
		t.Fatalf("Update returned nil error")
	}
	if !errors.Is(err, prepareErr) {
		t.Fatalf("Update error does not wrap prepare error: %v", err)
	}
}

func TestUpdateExecErrorWrapsCause(t *testing.T) {
	updateErr := errors.New("update failed")
	repo := stubRepository{
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, updateErr
		},
	}

	_, err := Update[queryTestModel](context.Background(), repo, map[string]any{"name": "alice"}, map[string]any{"name": "bob"})
	if err == nil {
		t.Fatalf("Update returned nil error")
	}
	if !errors.Is(err, updateErr) {
		t.Fatalf("Update error does not wrap exec error: %v", err)
	}
}

func TestUpdateRejectsTautologySqlizerPredicates(t *testing.T) {
	tests := []struct {
		name string
		expr sq.Sqlizer
	}{
		{name: "spaced", expr: sq.Expr("1 = 1")},
		{name: "parenthesized", expr: sq.Expr("(1 = 1)")},
		{name: "composed", expr: sq.And{sq.Expr("1 = 1")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := stubRepository{
				exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
					t.Fatalf("Exec should not be called")
					return pgconn.CommandTag{}, nil
				},
			}

			updated, err := Update[queryTestModel](context.Background(), repo, nil, map[string]any{"name": "bob"}, tt.expr)
			if err == nil {
				t.Fatalf("Update returned nil error")
			}
			if updated {
				t.Fatalf("Update returned true, want false")
			}
			if !strings.Contains(err.Error(), ErrEmptyPredicate) {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestUpdateRejectsEmptyPredicate(t *testing.T) {
	repo := stubRepository{
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			t.Fatalf("Exec should not be called")
			return pgconn.CommandTag{}, nil
		},
	}

	updated, err := Update[queryTestModel](context.Background(), repo, nil, map[string]any{"name": "bob"})
	if err == nil {
		t.Fatalf("Update returned nil error")
	}
	if updated {
		t.Fatalf("Update returned true, want false")
	}
	if !strings.Contains(err.Error(), ErrEmptyPredicate) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLogPubSubMsgUnexpectedCommandTagDoesNotPanic(t *testing.T) {
	repo := stubRepository{
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("INSERT 0 0"), nil
		},
	}

	err := LogPubSubMsg(context.Background(), repo, "topic", &pclient.Message{Data: []byte("{}")})
	if err == nil {
		t.Fatalf("LogPubSubMsg returned nil error")
	}
	if !strings.Contains(err.Error(), "pubsub_logs not created") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateRejectsEmptySqlizerPredicate(t *testing.T) {
	repo := stubRepository{
		exec: func(_ context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
			t.Fatalf("Exec should not be called")
			return pgconn.CommandTag{}, nil
		},
	}

	updated, err := Update[queryTestModel](context.Background(), repo, nil, map[string]any{"name": "bob"}, sq.Eq{})
	if err == nil {
		t.Fatalf("Update returned nil error")
	}
	if updated {
		t.Fatalf("Update returned true, want false")
	}
	if !strings.Contains(err.Error(), ErrEmptyPredicate) {
		t.Fatalf("unexpected error: %v", err)
	}
}
