package distributions

import (
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

// DistributionDistributionReport is an object representing the database table.
type DistributionDistributionReport struct {
	ID           int                 `db:"id" json:"id" yaml:"id"`
	UserID       *int                `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	OfferID      *int                `db:"offer_id" json:"offer_id,omitempty" yaml:"offer_id,omitempty"`
	Amount       float64             `db:"amount" json:"amount" yaml:"amount"`
	SubmitedAt   pgtype.Timestamptz  `db:"submited_at" json:"submited_at" yaml:"submited_at"`
	Data         any                 `db:"data" json:"data" yaml:"data"`
	CreatedAt    pgtype.Timestamptz  `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt    pgtype.Timestamptz  `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	ImageLinkID  *int                `db:"image_link_id" json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`
	Status       DistributionReportT `db:"status" json:"status" yaml:"status"`
	ApprovedByID *int                `db:"approved_by_id" json:"approved_by_id,omitempty" yaml:"approved_by_id,omitempty"`
	ApprovedAt   pgtype.Timestamptz  `db:"approved_at" json:"approved_at,omitempty" yaml:"approved_at,omitempty"`
}

func (model DistributionDistributionReport) ToJSON() map[string]any {
	return map[string]any{
		"id":             model.ID,
		"user_id":        model.UserID,
		"offer_id":       model.OfferID,
		"amount":         model.Amount,
		"submited_at":    model.SubmitedAt,
		"data":           model.Data,
		"created_at":     model.CreatedAt,
		"updated_at":     model.UpdatedAt,
		"image_link_id":  model.ImageLinkID,
		"status":         model.Status,
		"approved_by_id": model.ApprovedByID,
		"approved_at":    model.ApprovedAt,
	}
}

func (model DistributionDistributionReport) Fields() []string {
	return []string{
		"id", "user_id", "offer_id", "amount", "submited_at", "data",
		"created_at", "updated_at", "image_link_id",
	}
}

func (model DistributionDistributionReport) Table() string {
	return "distribution_distribution_reports"
}

func (model DistributionDistributionReport) GetID() any {
	return model.ID
}

func (model *DistributionDistributionReport) SetID(id any) {
	model.ID = id.(int)
}
