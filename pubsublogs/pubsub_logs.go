package pubsublogs

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/orm"
	"github.com/webdevelop-pro/go-common/orm/pgtype"
)

// PubsubLog is an object representing the database table.
type PubsubLog struct {
	ID        int                `db:"id" json:"id" yaml:"id"`
	Topic     string             `db:"topic" json:"topic" yaml:"topic"`
	MSG       any                `db:"msg" json:"msg" yaml:"msg"`
	Attr      any                `db:"attr" json:"attr,omitempty" yaml:"attr,omitempty"`
	MSGID     *string            `db:"msg_id" json:"msg_id,omitempty" yaml:"msg_id,omitempty"`
	Executed  int                `db:"executed" json:"executed,omitempty" yaml:"executed,omitempty"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	db        db.Repository      `db:"-" json:"-"`
}

func (model PubsubLog) ToJSON() map[string]any {
	return map[string]any{
		"id":         model.ID,
		"topic":      model.Topic,
		"msg":        model.MSG,
		"attr":       model.Attr,
		"msg_id":     model.MSGID,
		"executed":   model.Executed,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
	}
}

func (model PubsubLog) Fields() []string {
	return orm.DefaultFields(&model)
}

func (model PubsubLog) Table() string {
	return "pubsub_logs"
}

func (model PubsubLog) GetID() any {
	return model.ID
}

func (model *PubsubLog) SetID(id any) {
	model.ID = id.(int)
}

func (model *PubsubLog) SetDB(db db.Repository) {
	model.db = db
}

func Create(ctx context.Context, pg db.Repository, topic string, msg any, attr any, msgID string) (*PubsubLog, error) {
	return orm.Create[PubsubLog](ctx, pg, map[string]any{
		"topic":  topic,
		"msg":    msg,
		"attr":   attr,
		"msg_id": msgID,
	})
}

func IncrementExecuted(ctx context.Context, pg db.Repository, msgID string) (bool, error) {
	sql, args, err := sq.Update(PubsubLog{}.Table()).
		Set("executed", sq.Expr("executed+1")).
		Where(sq.Eq{"msg_id": msgID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, err
	}

	res, err := pg.Exec(ctx, sql, args...)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() == 1, nil
}
