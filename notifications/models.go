package notifications

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// NotificationNotification is an object representing the database table.
type NotificationNotification struct {
	ID        int                 `json:"id" yaml:"id"`
	UserID    int                 `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Content   string              `json:"content" yaml:"content"`
	Status    NotificationStatusT `json:"status" yaml:"status"`
	Type      NotificationTypeT   `json:"type" yaml:"type"`
	Data      any                 `json:"data" yaml:"data"`
	CreatedAt pgtype.Timestamptz  `json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz  `json:"updated_at" yaml:"updated_at"`
}

func (model NotificationNotification) ToJSON() map[string]any {
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

func (model NotificationNotification) Fields() []string {
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

func (model NotificationNotification) Table() string {
	return "notification_notifications"
}

func (model NotificationNotification) GetID() any {
	return model.ID
}

type Notifications []NotificationNotification
