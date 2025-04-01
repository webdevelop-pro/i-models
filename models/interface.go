package models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/webdevelop-pro/go-common/logger"
)

type Repository interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	// LogPubSubMsg(ctx context.Context, topic string, msg *pclient.Message) error
	Lg() logger.Logger
}

type Model interface {
	SetID(any)
	GetID() any
	ToJSON() map[string]any
	Fields() []string
	Table() string
}
