package notifications

import (
	"github.com/webdevelop-pro/i-models/pgtype"

	"github.com/webdevelop-pro/go-common/db"
)

// Notification is an object representing the database table.
type Notification struct {
	ID        int                 `json:"id" yaml:"id"`
	UserID    int                 `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Content   string              `json:"content" yaml:"content"`
	Status    NotificationStatusT `json:"status" yaml:"status"`
	Type      NotificationTypeT   `json:"type" yaml:"type"`
	Data      any                 `json:"data" yaml:"data"`
	CreatedAt pgtype.Timestamptz  `json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz  `json:"updated_at" yaml:"updated_at"`

	db db.Repository `db:"-" json:"-"`
}

func New(db db.Repository) *Notification {
	return &Notification{
		db: db,
	}
}

func (model Notification) ToJSON() map[string]any {
	return map[string]any{
		"id":         model.ID,
		"user_id":    model.UserID,
		"content":    model.Content,
		"status":     model.Status,
		"type":       model.Type,
		"data":       model.Data,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
	}
}

func (model Notification) Fields() []string {
	return []string{
		"ID",
		"UserID",
		"Content",
		"Status",
		"Type",
		"Data",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model Notification) Table() string {
	return "notification_notifications"
}

func (model Notification) GetID() any {
	return model.ID
}

func (model *Notification) SetID(id any) {
	model.ID = id.(int)
}

func (model *Notification) SetDB(db db.Repository) {
	model.db = db
}
