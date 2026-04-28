package emails

import (
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// EmailEmail is an object representing the database table.
type EmailEmail struct {
	ID                int                `json:"id" yaml:"id"`
	UserID            int                `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	TransactionID     *string            `json:"transaction_id,omitempty" yaml:"transaction_id,omitempty"`
	RecipientEmail    string             `json:"recipient_email" yaml:"recipient_email"`
	RecipientName     string             `json:"recipient_name" yaml:"recipient_name"`
	SenderEmail       string             `json:"sender_email" yaml:"sender_email"`
	SenderName        string             `json:"sender_name" yaml:"sender_name"`
	Subject           string             `json:"subject" yaml:"subject"`
	Template          string             `json:"template" yaml:"template"`
	Status            EmailStatusT       `json:"status" yaml:"status"`
	Data              any                `json:"data" yaml:"data"`
	SentAttemptCount  int                `json:"sent_attempt_count" yaml:"sent_attempt_count"`
	ContentHTML       string             `json:"content_html" yaml:"content_html"`
	MetaData          any                `json:"meta_data" yaml:"meta_data"`
	RecipientLocation any                `json:"recipient_location" yaml:"recipient_location"`
	Log               any                `json:"log" yaml:"log"`
	CreatedAt         pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
	SentAttemptAt     pgtype.Timestamptz `json:"sent_attempt_at" yaml:"sent_attempt_at"`

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
