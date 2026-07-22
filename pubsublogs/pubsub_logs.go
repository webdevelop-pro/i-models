package pubsublogs

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

// PubsubLog is an object representing the database table.
type PubsubLog struct {
	ID        int                `db:"id" json:"id" yaml:"id"`
	Topic     string             `db:"topic" json:"topic" yaml:"topic"`
	MSG       any                `db:"msg" json:"msg" yaml:"msg"`
	Headers   any                `db:"headers" json:"headers,omitempty" yaml:"headers,omitempty"`
	Attr      any                `db:"attr" json:"attr,omitempty" yaml:"attr,omitempty"`
	MSGID     *string            `db:"msg_id" json:"msg_id,omitempty" yaml:"msg_id,omitempty"`
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
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
	}
}

func (model PubsubLog) Fields() []string {
	return []string{
		"id", "topic", "msg", "attr", "msg_id", "created_at", "updated_at",
	}
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
