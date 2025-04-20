package historylogs

import (
	"time"
)

const Table = "django_admin_log"
const pkgName = "models/historylogs"

// LogLog is an object representing the database table.
type HistoryLog struct {
	ID            int       `json:"id" yaml:"id"`
	ContentTypeID int       `json:"content_type_id" yaml:"content_type_id"`
	ObjectID      string    `json:"object_id" yaml:"object_id"`
	ActionFlag    int       `json:"action_flag" yaml:"action_flag"`
	ObjectRepr    string    `json:"object_repr" yaml:"object_repr"`
	ChangeMessage string    `json:"change_message" yaml:"change_message"`
	UserID        int       `json:"user_id" yaml:"user_id"`
	ActionTime    time.Time `json:"action_time" yaml:"action_time"`
}

func (model HistoryLog) ToMap() map[string]any {
	return map[string]any{
		"id":              model.ID,
		"content_type_id": model.ContentTypeID,
		"object_id":       model.ObjectID,
	}
}

func (model HistoryLog) Fields() []string {
	// ToDo
	// Fix
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

func (model HistoryLog) Table() string {
	return Table
}

func (model HistoryLog) GetID() any {
	return model.ID
}

func (model *HistoryLog) SetID(id any) {
	model.ID = id.(int)
}
