package models

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/queue/pclient"
)

type queryModel[T any] interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}

func RetrieveOne[T any, PT queryModel[T]](ctx context.Context, pg Repository, where sq.Sqlizer, suffixes ...sq.Sqlizer) (*T, error) {
	obj := PT(new(T))

	builder := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table())
	if where != nil {
		builder = builder.Where(where)
	}
	for _, suffix := range suffixes {
		if suffix != nil {
			builder = builder.SuffixExpr(suffix)
		}
	}

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrSQLPrepare, where, suffixes,
		)
		return obj, resErr
	}

	rows, err := pg.Query(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrRetrieveOne, sql, args,
		)
		return obj, resErr
	}

	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[T])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &results, errors.Wrapf(
				ErrRecordNotFound,
				"%s: %s, %+v", ErrRetrieveOne, sql, args,
			)
		}
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrRetrieveOne, sql, args,
		)
		return &results, resErr
	}

	return &results, nil
}

// Deprecated: use RetrieveOne.
func RetriveOne[T any, PT queryModel[T]](ctx context.Context, pg Repository, where sq.Sqlizer, suffixes ...sq.Sqlizer) (*T, error) {
	return RetrieveOne[T, PT](ctx, pg, where, suffixes...)
}

func RetrieveAll[T any, PT queryModel[T]](ctx context.Context, pg Repository, where sq.Sqlizer, suffixes ...sq.Sqlizer) ([]*T, error) {
	obj := PT(new(T))

	builder := sq.Select(strings.Join(obj.Fields(), ",")).From(obj.Table())
	if where != nil {
		builder = builder.Where(where)
	}
	for _, suffix := range suffixes {
		if suffix != nil {
			builder = builder.SuffixExpr(suffix)
		}
	}

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrSQLPrepare, where, suffixes,
		)
		return nil, resErr
	}

	rows, err := pg.Query(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrRetrieveAll, sql, args,
		)
		return nil, resErr
	}

	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	results, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrRetrieveAll, sql, args,
		)
		return results, resErr
	}
	if results == nil {
		return []*T{}, nil
	}
	return results, nil
}

// Deprecated: use RetrieveAll.
func RetriveAll[T any, PT queryModel[T]](ctx context.Context, pg Repository, where sq.Sqlizer, suffixes ...sq.Sqlizer) ([]*T, error) {
	return RetrieveAll[T, PT](ctx, pg, where, suffixes...)
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
			err,
			"%s: %s, %+v", ErrSQLPrepare, sql, args,
		)
		return obj, resErr
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrCreate, sql, args,
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
}](ctx context.Context, pg db.Repository, where map[string]any, data map[string]any, exprs ...sq.Sqlizer) (bool, error) {
	obj := PT(new(T))
	if !hasPredicate(where, exprs) {
		return false, errors.Wrap(errors.New(ErrSQLPrepare), ErrEmptyPredicate)
	}

	builder := sq.Update(obj.Table()).SetMap(data)
	if len(where) > 0 {
		builder = builder.Where(where)
	}
	for _, expr := range exprs {
		if expr != nil {
			builder = builder.Where(expr)
		}
	}

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s %+v %+v", ErrSQLPrepare, where, data, exprs,
		)
		return false, resErr
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrUpdate, sql, args,
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
}](ctx context.Context, pg Repository, where map[string]any, exprs ...sq.Sqlizer) (bool, error) {
	obj := PT(new(T))
	res := 0

	builder := sq.Select("1").From(obj.Table())
	if len(where) > 0 {
		builder = builder.Where(where)
	}
	for _, expr := range exprs {
		if expr != nil {
			builder = builder.Where(expr)
		}
	}

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrSQLPrepare, where, exprs,
		)
		return false, resErr
	}

	err = pg.QueryRow(ctx, sql, args...).Scan(&res)
	// Assumes the returned row only has a single hit. StructToFill is the target struct.
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		} else {
			resErr := errors.Wrapf(
				err,
				"%s: %s, %+v", ErrRetrieveOne, sql, args,
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
}](ctx context.Context, pg Repository, where map[string]any, exprs ...sq.Sqlizer) (bool, error) {
	obj := PT(new(T))
	if !hasPredicate(where, exprs) {
		return false, errors.Wrap(errors.New(ErrSQLPrepare), ErrEmptyPredicate)
	}

	builder := sq.Delete(obj.Table())
	if len(where) > 0 {
		builder = builder.Where(where)
	}
	for _, expr := range exprs {
		if expr != nil {
			builder = builder.Where(expr)
		}
	}

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrSQLPrepare, where, exprs,
		)
		return false, resErr
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		resErr := errors.Wrapf(
			err,
			"%s: %s, %+v", ErrDelete, sql, args,
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

func hasPredicate(where map[string]any, exprs []sq.Sqlizer) bool {
	if len(where) > 0 {
		return true
	}
	for _, expr := range exprs {
		if isSubstantivePredicate(expr) {
			return true
		}
	}
	return false
}

func isSubstantivePredicate(expr sq.Sqlizer) bool {
	if expr == nil {
		return false
	}

	sql, _, err := expr.ToSql()
	if err != nil {
		return true
	}
	sql = strings.ToLower(strings.Join(strings.Fields(strings.TrimSpace(sql)), " "))
	if sql == "" {
		return false
	}

	stripped := trimOuterParens(sql)
	compactStripped := strings.ReplaceAll(stripped, " ", "")
	if stripped == "" || compactStripped == "1=1" || compactStripped == "true" {
		return false
	}

	compact := strings.ReplaceAll(sql, " ", "")
	return compact != "1=1" && compact != "true" &&
		!strings.Contains(compact, "(1=1)") && !strings.Contains(compact, "(true)")
}

func trimOuterParens(sql string) string {
	for len(sql) >= 2 && sql[0] == '(' && sql[len(sql)-1] == ')' && outerParensWrap(sql) {
		sql = strings.TrimSpace(sql[1 : len(sql)-1])
	}
	return sql
}

func outerParensWrap(sql string) bool {
	depth := 0
	for i, r := range sql {
		switch r {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 && i != len(sql)-1 {
				return false
			}
		}
		if depth < 0 {
			return false
		}
	}
	return depth == 0
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
			"%s %+v", topic, msg,
		)
		return resErr
	}
	return nil
}
