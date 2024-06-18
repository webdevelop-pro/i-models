package notifications

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

type Notification struct {
	ID        int64              `json:"id"`
	UserID    int                `json:"user_id"`
	Content   string             `json:"content"`
	Data      interface{}        `json:"data"`
	Status    StatusType         `json:"status"`
	Type      Type               `json:"type"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Notifications []Notification
