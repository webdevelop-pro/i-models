package models

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2"
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

func RetrieveOne[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) (*T, error) {
	obj, err := orm.RetrieveOne[T, PT](ctx, repo, sq.Eq(where))
	setDBIfSupported(obj, repo)
	return obj, err
}

func RetrieveAll[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) ([]*T, error) {
	results, err := orm.RetrieveAll[T, PT](ctx, repo, sq.Eq(where))
	for _, item := range results {
		setDBIfSupported(item, repo)
	}
	return results, err
}

// RetriveOne is retained for downstream source compatibility.
// Deprecated: use RetrieveOne.
func RetriveOne[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) (*T, error) {
	return RetrieveOne[T, PT](ctx, repo, where)
}

// RetriveAll is retained for downstream source compatibility.
// Deprecated: use RetrieveAll.
func RetriveAll[T any, PT queryModel[T]](ctx context.Context, repo db.Repository, where map[string]any) ([]*T, error) {
	return RetrieveAll[T, PT](ctx, repo, where)
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
