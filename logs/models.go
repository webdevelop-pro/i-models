package logs

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// LogLog is an object representing the database table.
type LogLog struct {
	ID               int                `json:"id" yaml:"id"`
	ContentTypeID    int                `json:"content_type_id" yaml:"content_type_id"`
	MSGID            *string            `json:"msg_id,omitempty" yaml:"msg_id,omitempty"`
	ObjectID         string             `json:"object_id" yaml:"object_id"`
	StatusCode       int                `json:"status_code" yaml:"status_code"`
	Path             string             `json:"path" yaml:"path"`
	RequestHeaders   any                `json:"request_headers" yaml:"request_headers"`
	RequestData      string             `json:"request_data" yaml:"request_data"`
	ResponseHeaders  any                `json:"response_headers" yaml:"response_headers"`
	ResponseData     string             `json:"response_data" yaml:"response_data"`
	Service          ServicesT          `json:"service" yaml:"service"`
	Type             LogTypeT           `json:"type" yaml:"type"`
	RequestID        *string            `json:"request_id,omitempty" yaml:"request_id,omitempty"`
	MetaData         any                `json:"meta_data" yaml:"meta_data"`
	RequestCreatedAt pgtype.Timestamptz `json:"request_created_at" yaml:"request_created_at"`
	CreatedAt        pgtype.Timestamptz `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt        pgtype.Timestamptz `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

func (model LogLog) ToJSON() map[string]any {
	return map[string]any{
		"id":                 model.ID,
		"content_type_id":    model.ContentTypeID,
		"msg_id":             model.MSGID,
		"object_id":          model.ObjectID,
		"status_code":        model.StatusCode,
		"path":               model.Path,
		"request_headers":    model.RequestHeaders,
		"request_data":       model.RequestData,
		"response_headers":   model.ResponseHeaders,
		"response_data":      model.ResponseData,
		"service":            model.Service,
		"type":               model.Type,
		"request_id":         model.RequestID,
		"meta_data":          model.MetaData,
		"request_created_at": model.RequestCreatedAt,
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
	}
}

func (model LogLog) Fields() []string {
	return []string{
		"ID",
		"ContentTypeID",
		"MSGID",
		"ObjectID",
		"StatusCode",
		"Path",
		"RequestHeaders",
		"RequestData",
		"ResponseHeaders",
		"ResponseData",
		"Service",
		"Type",
		"RequestID",
		"MetaData",
		"RequestCreatedAt",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model LogLog) Table() string {
	return "log_logs"
}

func (model LogLog) GetID() any {
	return model.ID
}
