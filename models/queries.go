package models

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
)

const pkgName = "models"

func log(ctx context.Context) *zerolog.Logger {
	return logger.FromCtx(ctx, pkgName)
}

func RetriveOne[T Model](ctx context.Context, pg Repository, where map[string]interface{}) (*T, error) {
	obj := *new(T)

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		// we need this to have stacktrace
		log(ctx).Error().Stack().Err(errors.WithStack(err)).Msg(ErrRetrieveOne)

		return nil, err
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		// we need this to have stacktrace
		log(ctx).Error().Stack().Err(errors.WithStack(err)).Msg(ErrRetrieveOne)

		return results, err
	}

	return results, nil
}

func RetriveAll[T Model](ctx context.Context, pg Repository, where map[string]interface{}) ([]*T, error) {
	obj := *new(T)

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		log(ctx).Error().Stack().Err(errors.New(ErrSQLPrepare)).Msg(ErrRetrieveAll)

		return nil, err
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		log(ctx).Error().Stack().Err(errors.New(ErrSQLRequest)).Msg(ErrRetrieveAll)

		return results, err
	}

	return results, nil
}

func Create[T any](ctx context.Context, pg Repository, obj Model) (int, error) {
	var (
		id   int
		data = obj.ToJSON()
		sqEq = sq.Eq{}
	)

	delete(data, "id")

	for key, val := range data {
		sqEq[key] = val
	}

	b := sq.Insert(obj.Table()).Suffix("RETURNING id").SetMap(sqEq)

	sql, args, err := b.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(errors.New(ErrSQLRequest), "%s: %s, %v", err.Error(), sql, args)
		log(ctx).Error().Stack().Err(errors.New(ErrSQLPrepare)).Msg(ErrCreate)

		return id, resErr
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		resErr := errors.Wrapf(err, "sql %s", db.CleanSQL(sql))
		log(ctx).Error().Stack().Err(errors.New(ErrSQLPrepare)).Msg(ErrCreate)

		return id, resErr
	}

	return id, nil
}
