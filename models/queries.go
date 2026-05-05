package models

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/queue/pclient"
)

func RetriveOne[T any, PT interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}](ctx context.Context, pg Repository, where map[string]any, queryParams ...string) (*T, error) {
	obj := PT(new(T))

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrSQLPrepare),
			"%s: %s, %+v", err.Error(), where, queryParams,
		)
		return obj, resErr
	}

	for _, param := range queryParams {
		sql = fmt.Sprintf("%s %s", sql, param)
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[T])
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrRetrieveOne),
			"%s: %s, %+v", err.Error(), sql, args,
		)
		return &results, resErr
	}

	return &results, nil
}

func RetriveAll[T any, PT interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}](ctx context.Context, pg db.Repository, where map[string]any, queryParams ...string) ([]*T, error) {
	obj := PT(new(T))

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrSQLPrepare),
			"%s: %s, %+v", err.Error(), where, queryParams,
		)
		return nil, resErr
	}

	for _, param := range queryParams {
		sql = fmt.Sprintf("%s %s", sql, param)
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrRetrieveAll),
			"%s: %s, %+v", err.Error(), sql, args,
		)
		return results, resErr
	}
	if results == nil {
		return []*T{}, nil
	}
	return results, nil
}

func Create[T any, PT interface {
	*T
	SetID(any)
	SetDB(db.Repository)
	Fields() []string
	Table() string
}](ctx context.Context, pg db.Repository, data map[string]any) (*T, error) {
	obj := PT(new(T))
	id := 0
	b := sq.Insert(obj.Table()).Suffix("RETURNING id").SetMap(data)
	sql, args, err := b.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrSQLPrepare),
			"%s: %s, %+v", err.Error(), sql, args,
		)
		return obj, resErr
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrCreate),
			"%s: %s, %+v", err.Error(), sql, args,
		)
		return obj, resErr
	}
	obj.SetID(id)
	obj.SetDB(pg)
	return obj, nil
}

func Update[T any, PT interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}](ctx context.Context, pg db.Repository, where map[string]any, data map[string]any) (bool, error) {
	obj := PT(new(T))
	b := sq.Update(obj.Table()).SetMap(data).Where(where)
	sql, args, err := b.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrSQLPrepare),
			"%s: %s %+v", err.Error(), where, data,
		)
		return false, resErr
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrUpdate),
			"%s: %s, %+v", err.Error(), sql, args,
		)
		return false, resErr
	}
	/*
		if res.String() != "UPDATE 1" {
			resErr := errors.Wrapf(errors.New(ErrSQLRequest), "no rows updated %s %+v", topic, msg)
			log.Error().Stack().Err(resErr).Msg("no rows updated")
			return resErr
		}
	*/
	return res.String() == "UPDATE 1", nil
}

func Exists[T any, PT interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}](ctx context.Context, pg Repository, where map[string]any) (bool, error) {
	obj := PT(new(T))
	res := 0

	sql, args, err := sq.Select("1").From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrSQLPrepare),
			"%s: %s", err.Error(), where,
		)
		return false, resErr
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&res)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return false, nil
		} else {
			resErr := errors.Wrapf(
				errors.New(ErrRetrieveOne),
				"%s: %s, %+v", err.Error(), sql, args,
			)
			return res == 1, resErr
		}
	}
	return res == 1, nil
}

func Delete[T any, PT interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}](ctx context.Context, pg Repository, where map[string]any) (bool, error) {
	obj := PT(new(T))

	sql, args, err := sq.Delete(obj.Table()).Where(where).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrSQLPrepare),
			"%s: %s", err.Error(), where,
		)
		return false, resErr
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(
			errors.New(ErrDelete),
			"%s: %s, %+v", err.Error(), sql, args,
		)
		return false, resErr
	}
	/*
		if res.String() != "UPDATE 1" {
			resErr := errors.Wrapf(errors.New(ErrSQLRequest), "no rows updated %s %+v", topic, msg)
			log.Error().Stack().Err(resErr).Msg("no rows updated")
			return resErr
		}
	*/
	return res.String() == "DELETE 1", nil
}

func LogPubSubMessageExecution(ctx context.Context, pg db.Repository, msgID string) error {
	query := `update pubsub_logs set executed=executed+1 where msg_id=$1;`

	_, err := pg.Exec(
		ctx,
		query,
		msgID,
	)
	if err != nil {
		resErr := errors.Wrapf(
			errors.New("UpdateInvestmentFundingStatus error"),
			"%s", err.Error(),
		)
		return resErr
	}

	return nil
}

func LogPubSubMsg(ctx context.Context, pg Repository, topic string, msg *pclient.Message) error {
	// ToDo
	// Use loki as a storage for pubsub logs insted of postgres
	sql := "INSERT INTO pubsub_logs (topic,msg,attr,headers,created_at,msg_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
	args := []interface{}{
		topic, msg.Data, msg.Attributes, msg.Headers, msg.PublishTime, msg.ID,
	}

	// for some reason we don't always get ID from google pubsub
	if msg.ID == "" {
		args[5] = nil
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(
			errors.New("UpdateInvestmentFundingStatus error"),
			"%s, %s %+v", err.Error(), sql, args,
		)
		return resErr
	}
	if res.String() != "INSERT 0 1" { // event sequence haven't been updated
		resErr := errors.Wrapf(
			errors.New("pubsub_logs not created"),
			"%s, %s %+v", err.Error(), topic, msg,
		)
		return resErr
	}
	return nil
}
