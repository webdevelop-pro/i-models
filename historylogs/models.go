package historylogs

import (
	"time"

	"github.com/global-torque/go-common/db/v2"
)

const Table = "django_admin_log"
const pkgName = "models/historylogs"

// LogLog is an object representing the database table.
type HistoryLog struct {
	ID            int       `db:"id" json:"id" yaml:"id"`
	ContentTypeID *int      `db:"content_type_id" json:"content_type_id,omitempty" yaml:"content_type_id,omitempty"`
	ObjectID      *string   `db:"object_id" json:"object_id,omitempty" yaml:"object_id,omitempty"`
	ActionFlag    int       `db:"action_flag" json:"action_flag" yaml:"action_flag"`
	ObjectRepr    string    `db:"object_repr" json:"object_repr" yaml:"object_repr"`
	ChangeMessage string    `db:"change_message" json:"change_message" yaml:"change_message"`
	UserID        int       `db:"user_id" json:"user_id" yaml:"user_id"`
	ActionTime    time.Time `db:"action_time" json:"action_time" yaml:"action_time"`

	db db.Repository `db:"-" json:"-"`
}

func New(db db.Repository) *HistoryLog {
	return &HistoryLog{
		db: db,
	}
}

func (model HistoryLog) ToMap() map[string]any {
	return map[string]any{
		"id":              model.ID,
		"content_type_id": model.ContentTypeID,
		"object_id":       model.ObjectID,
	}
}

func (model HistoryLog) Fields() []string {
	return []string{
		"id", "content_type_id", "object_id", "action_flag", "object_repr",
		"change_message", "user_id", "action_time",
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

func (model *HistoryLog) SetDB(db db.Repository) {
	model.db = db
}
