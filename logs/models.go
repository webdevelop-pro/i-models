package logs

import "github.com/jackc/pgtype"

type Logs struct {
	ID              int            `json:"id"`
	ContentTypeID   int            `json:"content_type_id"`
	ObjectID        int            `json:"object_id"`
	StatusCode      string         `json:"status_code"`
	Path            string         `json:"path"`
	RequestHeaders  map[string]any `json:"request_headers"`
	RequestData     string         `json:"request_data"`
	ResponseHeaders map[string]any `json:"response_headers"`
	ResponseData    string         `json:"response_data"`
	Service         ServiceType    `json:"service"`
	Type            Type           `json:"type"`
	RequestID       string         `json:"request_id"`
	MetaData        map[string]any `json:"meta_data"`
	MsgID           string         `json:"msg_id"`

	RequestCreatedAt pgtype.Timestamptz `json:"request_created_at"`
	CreatedAt        pgtype.Timestamptz `json:"created_at"`
	UpdatedAt        pgtype.Timestamptz `json:"updated_at"`
}
