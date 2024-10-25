package offers

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// OfferComment is an object representing the database table.
type OfferComment struct {
	ID         int                 `json:"id" yaml:"id"`
	UserID     int                 `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	OfferID    int                 `json:"offer_id,omitempty" yaml:"offer_id,omitempty"`
	Comment    string              `json:"comment" yaml:"comment"`
	Related    NullCommentRelatedT `json:"related,omitempty" yaml:"related,omitempty"`
	Status     NullCommentStatusT  `json:"status,omitempty" yaml:"status,omitempty"`
	ApprovedAt pgtype.Timestamptz  `json:"approved_at,omitempty" yaml:"approved_at,omitempty"`
	CreatedAt  pgtype.Timestamptz  `json:"created_at" yaml:"created_at"`
	UpdatedAt  pgtype.Timestamptz  `json:"updated_at" yaml:"updated_at"`
}

func (model OfferComment) ToJSON() map[string]any {
	return map[string]any{
		"id":          model.ID,
		"user_id":     model.UserID,
		"offer_id":    model.OfferID,
		"comment":     model.Comment,
		"related":     model.Related,
		"status":      model.Status,
		"approved_at": model.ApprovedAt,
		"created_at":  model.CreatedAt,
		"updated_at":  model.UpdatedAt,
	}
}

func (model OfferComment) Fields() []string {
	return []string{
		"ID",
		"UserID",
		"OfferID",
		"Comment",
		"Related",
		"Status",
		"ApprovedAt",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model OfferComment) Table() string {
	return "offer_comments"
}

func (model OfferComment) GetID() any {
	return model.ID
}
