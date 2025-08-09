package models

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/go-common/queue/pclient"
)

func RetriveOne[T any, PT interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}](ctx context.Context, pg Repository, where map[string]any) (*T, error) {
	log := logger.FromCtx(ctx, "models")
	obj := PT(new(T))

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		// we need this to have stacktrace
		log.Error().Stack().Err(errors.WithStack(err)).Msg(ErrRetrieveOne)
		return obj, err
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	if err != nil {
		// we need this to have stacktrace
		log.Error().Stack().Err(errors.WithStack(err)).Msg(ErrRetrieveOne)
		return &results, err
	}

	return &results, nil
}

func RetriveAll[T any, PT interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}](ctx context.Context, pg db.Repository, where map[string]any) ([]*T, error) {
	log := logger.FromCtx(ctx, "models")
	obj := PT(new(T))

	// ToDo
	// Add where in the loop with where incoming parameter
	sql, args, err := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		log.Error().Stack().Err(err).Msg(ErrRetrieveAll)
		return nil, err
	}

	rows, _ := pg.Query(ctx, sql, args...)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		log.Error().Stack().Err(err).Msg(ErrRetrieveAll)
		return results, err
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
	log := logger.FromCtx(ctx, "models")
	obj := PT(new(T))
	id := 0
	b := sq.Insert(obj.Table()).Suffix("RETURNING id").SetMap(data)
	sql, args, err := b.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(errors.New(ErrSQLRequest), "%s: %s, %v", err.Error(), sql, args)
		log.Error().Stack().Err(resErr).Msg(ErrCreate)
		return obj, resErr
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		resErr := errors.Wrapf(err, "sql error %s", obj.Table())
		log.Error().Stack().Any("sql", sql).Any("data", data).Err(resErr).Msg(ErrCreate)
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
	log := logger.FromCtx(ctx, "models")
	obj := PT(new(T))
	b := sq.Update(obj.Table()).SetMap(data).Where(where)
	sql, args, err := b.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(errors.New(ErrSQLRequest), "%s: %s, %v", err.Error(), sql, args)
		log.Error().Stack().Err(resErr).Msg(ErrUpdate)
		return false, resErr
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(err, "sql %s", db.CleanSQL(sql))
		log.Error().Stack().Err(resErr).Msg(ErrUpdate)
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
	log := logger.FromCtx(ctx, "models")
	obj := PT(new(T))
	res := 0

	sql, args, err := sq.Select("1").From(obj.Table()).
		Where(where).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		// we need this to have stacktrace
		log.Error().Stack().Err(err).Msg(ErrRetrieveOne)
		return false, err
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&res)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return false, nil
		} else {
			log.Error().Stack().Err(err).Msg(ErrRetrieveOne)
			return res == 1, err
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
	log := logger.FromCtx(ctx, "models")
	obj := PT(new(T))

	sql, args, err := sq.Delete(obj.Table()).Where(where).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		// we need this to have stacktrace
		log.Error().Stack().Err(errors.WithStack(err)).Msg(ErrRetrieveOne)
		return false, err
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(err, "sql %s", db.CleanSQL(sql))
		log.Error().Stack().Err(resErr).Msg(ErrUpdate)
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
	log := logger.FromCtx(ctx, "models")

	_, err := pg.Exec(
		ctx,
		query,
		msgID,
	)
	if err != nil {
		err = errors.Wrapf(err, "for msg %s", msgID)
		log.Error().Stack().Err(err).Msg("UpdateInvestmentFundingStatus error")

		return err
	}

	return nil
}

func LogPubSubMsg(ctx context.Context, pg Repository, topic string, msg *pclient.Message) error {
	// ToDo
	// Use loki as a storage for pubsub logs insted of postgres
	log := logger.FromCtx(ctx, "models")
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
		resErr := errors.Wrapf(err, "sql %s, args: %+v", sql, args)
		log.Error().Stack().Err(resErr).Msg("LogPubSubMsg error")
		return resErr
	}
	if res.String() != "INSERT 0 1" { // event sequence haven't been updated
		resErr := errors.Wrapf(db.ErrSQLRequest, "pubsub_logs not created %s %+v", topic, msg)
		log.Error().Stack().Err(resErr).Msg("LogPubSubMsg error")
		return resErr
	}
	return nil
}
