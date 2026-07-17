package emails

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/google/uuid"
)

// EmailEmail is an object representing the database table.
type EmailEmail struct {
	ID                int                `db:"id" json:"id" yaml:"id"`
	UserID            *int               `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	TransactionID     *string            `db:"transaction_id" json:"transaction_id,omitempty" yaml:"transaction_id,omitempty"`
	RecipientEmail    string             `db:"recipient_email" json:"recipient_email" yaml:"recipient_email"`
	RecipientName     string             `db:"recipient_name" json:"recipient_name" yaml:"recipient_name"`
	SenderEmail       string             `db:"sender_email" json:"sender_email" yaml:"sender_email"`
	SenderName        string             `db:"sender_name" json:"sender_name" yaml:"sender_name"`
	Subject           string             `db:"subject" json:"subject" yaml:"subject"`
	Template          string             `db:"template" json:"template" yaml:"template"`
	Status            EmailStatusT       `db:"status" json:"status" yaml:"status"`
	Data              any                `db:"data" json:"data" yaml:"data"`
	SentAttemptCount  int                `db:"sent_attempt_count" json:"sent_attempt_count" yaml:"sent_attempt_count"`
	ContentHTML       string             `db:"content_html" json:"content_html" yaml:"content_html"`
	MetaData          any                `db:"meta_data" json:"meta_data" yaml:"meta_data"`
	RecipientLocation any                `db:"recipient_location" json:"recipient_location" yaml:"recipient_location"`
	Log               any                `db:"log" json:"log" yaml:"log"`
	CreatedAt         pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt         pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	SentAttemptAt     pgtype.Timestamptz `db:"sent_attempt_at" json:"sent_attempt_at" yaml:"sent_attempt_at"`
	DomainEventID     *uuid.UUID         `db:"domain_event_id" json:"domain_event_id,omitempty" yaml:"domain_event_id,omitempty"`

	db db.Repository `db:"-" json:"-"`
}

func (model EmailEmail) ToJSON() map[string]any {
	return map[string]any{
		"id":                 model.ID,
		"user_id":            model.UserID,
		"transaction_id":     model.TransactionID,
		"recipient_email":    model.RecipientEmail,
		"recipient_name":     model.RecipientName,
		"sender_email":       model.SenderEmail,
		"sender_name":        model.SenderName,
		"subject":            model.Subject,
		"template":           model.Template,
		"status":             model.Status,
		"data":               model.Data,
		"sent_attempt_count": model.SentAttemptCount,
		"content_html":       model.ContentHTML,
		"meta_data":          model.MetaData,
		"recipient_location": model.RecipientLocation,
		"log":                model.Log,
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
		"sent_attempt_at":    model.SentAttemptAt,
		"domain_event_id":    model.DomainEventID,
	}
}

func (model EmailEmail) Fields() []string {
	return []string{
		"id",
		"user_id",
		"transaction_id",
		"recipient_email",
		"recipient_name",
		"sender_email",
		"sender_name",
		"subject",
		"template",
		"status",
		"data",
		"sent_attempt_count",
		"content_html",
		"meta_data",
		"recipient_location",
		"log",
		"created_at",
		"updated_at",
		"sent_attempt_at",
	}
}

func (model EmailEmail) Table() string {
	return "email_emails"
}

func (model EmailEmail) GetID() any {
	return model.ID
}

func (model *EmailEmail) SetID(id any) {
	model.ID = id.(int)
}

func (model *EmailEmail) SetDB(db db.Repository) {
	model.db = db
}
