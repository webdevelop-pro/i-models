package models

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
	"github.com/jackc/pgx/v5"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
)

func RetriveOne[T Model](ctx context.Context, pg Repository, where map[string]interface{}) (*T, error) {
	log := logger.FromCtx(ctx, "models")
	obj := *new(T)

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		// we need this to have stacktrace
		log.Error().Stack().Err(errors.New(ErrSQLPrepare)).Msg(ErrRetrieveOne)
		return nil, err
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		// we need this to have stacktrace
		log.Error().Stack().Err(errors.New(ErrSQLRequest)).Msg(ErrRetrieveOne)
		return results, err
	}
	return results, nil
}

func RetriveAll[T Model](ctx context.Context, pg Repository, where map[string]interface{}) ([]*T, error) {
	log := logger.FromCtx(ctx, "models")
	obj := *new(T)

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		log.Error().Stack().Err(errors.New(ErrSQLPrepare)).Msg(ErrRetrieveAll)
		return nil, err
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		log.Error().Stack().Err(errors.New(ErrSQLRequest)).Msg(ErrRetrieveAll)
		return results, err
	}
	return results, nil
}

func Create[T any](ctx context.Context, pg Repository, obj Model) (int, error) {
	log := logger.FromCtx(ctx, "models")
	id := 0
	data := obj.ToJSON()
	delete(data, "id")
	sqEq := sq.Eq{}
	for key, val := range data {
		sqEq[key] = val
	}
	b := sq.Insert(obj.Table()).Suffix("RETURNING id").SetMap(sqEq)
	sql, args, err := b.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(errors.New(ErrSQLRequest), "%s: %s, %v", err.Error(), sql, args)
		log.Error().Stack().Err(errors.New(ErrSQLPrepare)).Msg(ErrCreate)
		return id, resErr
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		resErr := errors.Wrapf(err, "sql %s", db.CleanSQL(sql))
		log.Error().Stack().Err(errors.New(ErrSQLPrepare)).Msg(ErrCreate)
		return id, resErr
	}
	return id, nil
}
