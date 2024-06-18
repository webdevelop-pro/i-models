package pubsublogs

import "github.com/webdevelop-pro/i-models/pgtype"

type Logs struct {
	ID       int            `json:"id"`
	Topic    string         `json:"topic"`
	MsgID    string         `json:"msg_id"`
	Msg      any            `json:"msg"`
	Attr     map[string]any `json:"attr"`
	Executed int            `json:"executed"`

	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
