package distributions

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// DistributionDistributionReport is an object representing the database table.
type DistributionDistributionReport struct {
	ID          int                `json:"id" yaml:"id"`
	UserID      int                `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	OfferID     int                `json:"offer_id,omitempty" yaml:"offer_id,omitempty"`
	Amount      float64            `json:"amount" yaml:"amount"`
	SubmitedAt  pgtype.Timestamptz `json:"submited_at" yaml:"submited_at"`
	Data        any                `json:"data" yaml:"data"`
	CreatedAt   pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
	ImageLinkID int                `json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`
}

func (model DistributionDistributionReport) ToJSON() map[string]any {
	return map[string]any{
		"id":            model.ID,
		"user_id":       model.UserID,
		"offer_id":      model.OfferID,
		"amount":        model.Amount,
		"submited_at":   model.SubmitedAt,
		"data":          model.Data,
		"created_at":    model.CreatedAt,
		"updated_at":    model.UpdatedAt,
		"image_link_id": model.ImageLinkID,
	}
}

func (model DistributionDistributionReport) Fields() []string {
	return []string{
		"ID",
		"UserID",
		"OfferID",
		"Amount",
		"SubmitedAt",
		"Data",
		"CreatedAt",
		"UpdatedAt",
		"ImageLinkID",
	}
}

func (model DistributionDistributionReport) Table() string {
	return "distribution_distribution_reports"
}

func (model DistributionDistributionReport) GetID() any {
	return model.ID
}
