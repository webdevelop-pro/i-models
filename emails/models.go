package emails

import "github.com/jackc/pgtype"

type Email struct {
	ID               int     `json:"id"`
	Status           string  `json:"status"`
	UserID           *int    `json:"user_id"`
	TransactionID    *string `json:"transaction_id"`
	SentAttemptCount int     `json:"sent_attempt_count"`
	RecipientEmail   string  `json:"recipient_email"`
	RecipientName    string  `json:"recipient_name"`
	SenderEmail      string  `json:"sender_email"`
	SenderName       string  `json:"sender_name"`
	Subject          string  `json:"subject"`
	Template         string  `json:"template"`
	ContentText      string  `json:"-"`
	ContentHTML      string  `json:"-"`

	Data              interface{}            `json:"data"`
	MetaData          map[string]interface{} `json:"meta_data"`
	RecipientLocation map[string]interface{} `json:"recipient_location"`

	SentAttempAt pgtype.Timestamptz `json:"sent_attemp_at"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}
