package models

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/orm"
	"github.com/webdevelop-pro/go-common/queue/pclient"
)

type Repository = db.Repository

type queryModel[T any] interface {
	*T
	SetID(any)
	Fields() []string
	Table() string
}

type dbSetter interface {
	SetDB(db.Repository)
}

func setDBIfSupported(obj any, repo db.Repository) {
	if setter, ok := obj.(dbSetter); ok {
		setter.SetDB(repo)
	}
}

func DefaultFields(obj any) []string {
	return orm.DefaultFields(obj)
}

func RetriveOne[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) (*T, error) {
	obj, err := orm.RetrieveOne[T, PT](ctx, repo, sq.Eq(where))
	setDBIfSupported(obj, repo)
	return obj, err
}

func RetriveAll[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) ([]*T, error) {
	results, err := orm.RetrieveAll[T, PT](ctx, repo, sq.Eq(where))
	for _, item := range results {
		setDBIfSupported(item, repo)
	}
	return results, err
}

func Create[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, data map[string]any) (*T, error) {
	obj, err := orm.Create[T, PT](ctx, repo, data)
	setDBIfSupported(obj, repo)
	return obj, err
}

func Update[T any, PT queryModel[T]](
	ctx context.Context,
	repo db.Repository,
	where map[string]any,
	data map[string]any,
) (bool, error) {
	return orm.Update[T, PT](ctx, repo, where, data)
}

func Exists[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) (bool, error) {
	return orm.Exists[T, PT](ctx, repo, where)
}

func Delete[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) (bool, error) {
	return orm.Delete[T, PT](ctx, repo, where)
}

func LogPubSubMessageExecution(ctx context.Context, repo db.Repository, msgID string) error {
	_, err := repo.Exec(ctx, `UPDATE pubsub_logs SET executed=executed+1 WHERE msg_id=$1`, msgID)
	if err != nil {
		return fmt.Errorf("for msg %s: %w", msgID, err)
	}

	return nil
}

func LogPubSubMsg(ctx context.Context, repo db.Repository, topic string, msg *pclient.Message) error {
	msgID := any(msg.ID)
	if msg.ID == "" {
		msgID = nil
	}

	_, err := repo.Exec(
		ctx,
		`INSERT INTO pubsub_logs (topic,msg,attr,headers,created_at,msg_id) VALUES ($1,$2,$3,$4,$5,$6)`,
		topic,
		msg.Data,
		msg.Attributes,
		msg.Headers,
		msg.PublishTime,
		msgID,
	)
	if err != nil {
		return fmt.Errorf("log pubsub message %s: %w", topic, err)
	}

	return nil
}
