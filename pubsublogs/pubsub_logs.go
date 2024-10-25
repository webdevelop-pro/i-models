package pubsublogs

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// PubsubLog is an object representing the database table.
type PubsubLog struct {
	ID        int                `json:"id" yaml:"id"`
	Topic     string             `json:"topic" yaml:"topic"`
	MSG       any                `json:"msg" yaml:"msg"`
	Attr      any                `json:"attr,omitempty" yaml:"attr,omitempty"`
	MSGID     *string            `json:"msg_id,omitempty" yaml:"msg_id,omitempty"`
	Executed  int                `json:"executed,omitempty" yaml:"executed,omitempty"`
	CreatedAt pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
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
	return []string{
		"ID",
		"Topic",
		"MSG",
		"Attr",
		"MSGID",
		"Executed",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model PubsubLog) Table() string {
	return "pubsub_logs"
}

func (model PubsubLog) GetID() any {
	return model.ID
}
