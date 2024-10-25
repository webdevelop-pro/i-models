package distributions

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// DistributionDistribution is an object representing the database table.
type DistributionDistribution struct {
	ID           int                `json:"id" yaml:"id"`
	UserID       int                `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	ProfileID    int                `json:"profile_id,omitempty" yaml:"profile_id,omitempty"`
	InvestmentID int                `json:"investment_id,omitempty" yaml:"investment_id,omitempty"`
	ReportID     int                `json:"report_id,omitempty" yaml:"report_id,omitempty"`
	Amount       float64            `json:"amount" yaml:"amount"`
	Status       DistributionT      `json:"status" yaml:"status"`
	Data         any                `json:"data" yaml:"data"`
	SubmitedAt   pgtype.Timestamptz `json:"submited_at" yaml:"submited_at"`
	CreatedAt    pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
}

func (model DistributionDistribution) ToJSON() map[string]any {
	return map[string]any{
		"id":            model.ID,
		"user_id":       model.UserID,
		"profile_id":    model.ProfileID,
		"investment_id": model.InvestmentID,
		"report_id":     model.ReportID,
		"amount":        model.Amount,
		"status":        model.Status,
		"data":          model.Data,
		"submited_at":   model.SubmitedAt,
		"created_at":    model.CreatedAt,
		"updated_at":    model.UpdatedAt,
	}
}

func (model DistributionDistribution) Fields() []string {
	return []string{
		"ID",
		"UserID",
		"ProfileID",
		"InvestmentID",
		"ReportID",
		"Amount",
		"Status",
		"Data",
		"SubmitedAt",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model DistributionDistribution) Table() string {
	return "distribution_distributions"
}

func (model DistributionDistribution) GetID() any {
	return model.ID
}
