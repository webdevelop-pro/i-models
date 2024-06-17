package models

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
)

type Store[M Model] struct {
	Pool *db.DB
	//nolint:structcheck
	table string
	//nolint:structcheck
	item *M
}

func NewStore[M Model](pool *db.DB, table string, item M) *Store[M] {
	return &Store[M]{
		Pool:  pool,
		table: table,
		item:  &item,
	}
}

func (s *Store[M]) Item() *M {
	return s.item
}

func (s *Store[M]) Create(ctx context.Context, m M) error {
	//nolint:forbidigo
	fmt.Printf("%v %v, create model \n", ctx, m)
	return nil
}

func (s *Store[M]) fields() []string {
	vals := reflect.ValueOf(s.item).Elem()
	valsLen := vals.NumField()
	fields := make([]string, 0)

	for i := 0; i < valsLen; i++ {
		dbTag := vals.Type().Field(i).Tag.Get("db")
		if dbTag != "" {
			fields = append(fields, dbTag)
		}
	}
	return fields
}

func (s *Store[M]) Select(ctx context.Context, params any) (*M, error) {
	fields := s.fields()
	sql, args, err := sq.Select(strings.Join(fields, ",")).From(s.table).
		Where(params).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(err, "%s, %v", sql, args)
		return s.item, resErr
	}

	rows, _ := s.Pool.Pool.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[M])
	//nolint:gocritic
	if err != nil {
		s.Pool.Log.Error().Err(err).Msg("Select error")
		return s.item, err
	} else if len(res) > 0 {
		s.item = &res[0]
	} else {
		return s.item, pgx.ErrNoRows
	}
	return s.item, nil
}

func (s *Store[M]) UpdateByID(ctx context.Context, values map[string]any) (string, error) {
	builder := sq.Update(s.table)
	for key, val := range values {
		builder = builder.Set(key, val)
	}
	builder = builder.Where("id=?", (*s.item).PrimaryFieldValue())
	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(err, "%s, %v", sql, args)
		return "", resErr
	}
	res, err := s.Pool.Pool.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(err, "for id %d", (*s.item).PrimaryFieldValue())
		return res.String(), resErr
	}
	return res.String(), nil
}
