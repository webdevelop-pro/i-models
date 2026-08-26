package notifications

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/google/uuid"
)

// Notification is an object representing the database table.
type Notification struct {
	ID                   int                 `db:"id" json:"id" yaml:"id"`
	UserID               int                 `db:"user_id" json:"user_id" yaml:"user_id"`
	Content              string              `db:"content" json:"content" yaml:"content"`
	Status               NotificationStatusT `db:"status" json:"status" yaml:"status"`
	Type                 NotificationTypeT   `db:"type" json:"type" yaml:"type"`
	Data                 any                 `db:"data" json:"data" yaml:"data"`
	CreatedAt            pgtype.Timestamptz  `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt            pgtype.Timestamptz  `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	DomainEventID        *uuid.UUID          `db:"domain_event_id" json:"domain_event_id,omitempty" yaml:"domain_event_id,omitempty"`
	NotificationRevision int64               `db:"notification_revision" json:"notification_revision" yaml:"notification_revision"`
	NovuPushProcessedAt  pgtype.Timestamptz  `db:"novu_push_processed_at" json:"-" yaml:"-"`
	NovuPushAttemptCount int                 `db:"novu_push_attempt_count" json:"-" yaml:"-"`

	db db.Repository `db:"-" json:"-"`
}

func New(db db.Repository) *Notification {
	return &Notification{
		db: db,
	}
}

func (model Notification) ToJSON() map[string]any {
	return map[string]any{
		"id":                    model.ID,
		"user_id":               model.UserID,
		"content":               model.Content,
		"status":                model.Status,
		"type":                  model.Type,
		"data":                  model.Data,
		"created_at":            model.CreatedAt,
		"updated_at":            model.UpdatedAt,
		"domain_event_id":       model.DomainEventID,
		"notification_revision": model.NotificationRevision,
	}
}

func (model Notification) Fields() []string {
	return []string{
		"id",
		"user_id",
		"content",
		"status",
		"type",
		"data",
		"created_at",
		"updated_at",
		"notification_revision",
	}
}

func (model Notification) Table() string {
	return TableName
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
