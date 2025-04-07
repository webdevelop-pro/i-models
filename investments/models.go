package investments

import (
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// InvestmentInvestment is an object representing the database table.
type InvestmentInvestment struct {
	ID                int                `db:"id" json:"id" yaml:"id"`
	UserID            int                `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	OfferID           int                `db:"offer_id" json:"offer_id,omitempty" yaml:"offer_id,omitempty"`
	ProfileID         int                `db:"profile_id" json:"profile_id,omitempty" yaml:"profile_id,omitempty"`
	Amount            float64            `db:"amount" json:"amount" yaml:"amount"`
	PricePerShare     float64            `db:"price_per_share" json:"price_per_share" yaml:"price_per_share"`
	NumberOfShares    int                `db:"number_of_shares" json:"number_of_shares" yaml:"number_of_shares"`
	PaymentType       PaymentT           `db:"payment_type" json:"payment_type" yaml:"payment_type"`
	EscrowType        EscrowT            `db:"escrow_type" json:"escrow_type" yaml:"escrow_type"`
	FundingType       FundingT           `db:"funding_type" json:"funding_type" yaml:"funding_type"`
	FundingStatus     FundingS           `db:"funding_status" json:"funding_status" yaml:"funding_status"`
	Status            InvestmentT        `db:"status" json:"status" yaml:"status"`
	PrevStatus        InvestmentT        `db:"prev_status" json:"prev_status" yaml:"prev_status"`
	Step              InvestmentStepT    `db:"step" json:"step" yaml:"step"`
	Commission        float64            `db:"commission" json:"commission" yaml:"commission"`
	CancelationReason string             `db:"cancelation_reason" json:"cancelation_reason" yaml:"cancelation_reason"`
	EntityID          *string            `db:"entity_id" json:"entity_id,omitempty" yaml:"entity_id,omitempty"`
	TransactionRef    *string            `db:"transaction_ref" json:"transaction_ref,omitempty" yaml:"transaction_ref,omitempty"`
	EscrowData        any                `db:"escrow_data" json:"escrow_data" yaml:"escrow_data"`
	SignatureData     any                `db:"signature_data" json:"signature_data" yaml:"signature_data"`
	PaymentData       any                `db:"payment_data" json:"payment_data" yaml:"payment_data"`
	CanceledAt        pgtype.Timestamptz `db:"canceled_at" json:"canceled_at" yaml:"canceled_at"`
	SubmitedAt        pgtype.Timestamptz `db:"submited_at" json:"submited_at" yaml:"submited_at"`
	ClosedAt          pgtype.Timestamptz `db:"closed_at" json:"closed_at" yaml:"closed_at"`
	CreatedAt         pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt         pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`
}

// IsFundingTypeWire temporary fix for https://github.com/Joker/jade/issues/59
func (model InvestmentInvestment) IsFundingTypeWire() bool {
	return model.FundingType == "wire"
}

func (model InvestmentInvestment) ToJSON() map[string]any {
	return map[string]any{
		"id":                 model.ID,
		"user_id":            model.UserID,
		"offer_id":           model.OfferID,
		"profile_id":         model.ProfileID,
		"amount":             model.Amount,
		"price_per_share":    model.PricePerShare,
		"number_of_shares":   model.NumberOfShares,
		"payment_type":       model.PaymentType,
		"escrow_type":        model.EscrowType,
		"funding_type":       model.FundingType,
		"funding_status":     model.FundingStatus,
		"status":             model.Status,
		"prev_status":        model.PrevStatus,
		"step":               model.Step,
		"commission":         model.Commission,
		"cancelation_reason": model.CancelationReason,
		"entity_id":          model.EntityID,
		"transaction_ref":    model.TransactionRef,
		"escrow_data":        model.EscrowData,
		"signature_data":     model.SignatureData,
		"payment_data":       model.PaymentData,
		"canceled_at":        model.CanceledAt,
		"submited_at":        model.SubmitedAt,
		"closed_at":          model.ClosedAt,
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
	}
}

func (model InvestmentInvestment) Fields() []string {
	return models.DefaultFields(&model)
}

func (model InvestmentInvestment) Table() string {
	return "investment_investments"
}

func (model InvestmentInvestment) GetID() any {
	return model.ID
}

func (model *InvestmentInvestment) SetID(id any) {
	model.ID = id.(int)
}

// InvestmentProfile is an object representing the database table.
type InvestmentProfile struct {
	ID                  int                `json:"id" yaml:"id"`
	UserID              int                `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Type                ProfileT           `json:"type" yaml:"type"`
	Data                any                `json:"data" yaml:"data"`
	KycID               *string            `json:"kyc_id,omitempty" yaml:"kyc_id,omitempty"`
	KycStatus           KycT               `json:"kyc_status" yaml:"kyc_status"`
	KycData             any                `json:"kyc_data" yaml:"kyc_data"`
	AccreditationID     *string            `json:"accreditation_id,omitempty" yaml:"accreditation_id,omitempty"`
	AccreditationStatus AccreditationT     `json:"accreditation_status" yaml:"accreditation_status"`
	AccreditationData   any                `json:"accreditation_data" yaml:"accreditation_data"`
	EscrowID            string             `json:"escrow_id" yaml:"escrow_id"`
	KycAt               pgtype.Timestamptz `json:"kyc_at,omitempty" yaml:"kyc_at,omitempty"`
	AccreditationAt     pgtype.Timestamptz `json:"accreditation_at,omitempty" yaml:"accreditation_at,omitempty"`
	CreatedAt           pgtype.Timestamptz `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt           pgtype.Timestamptz `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	WalletID            int                `json:"wallet_id,omitempty" yaml:"wallet_id,omitempty"`
}

func (model InvestmentProfile) ToJSON() map[string]any {
	return map[string]any{
		"id":                   model.ID,
		"user_id":              model.UserID,
		"type":                 model.Type,
		"data":                 model.Data,
		"kyc_id":               model.KycID,
		"kyc_status":           model.KycStatus,
		"kyc_data":             model.KycData,
		"accreditation_id":     model.AccreditationID,
		"accreditation_status": model.AccreditationStatus,
		"accreditation_data":   model.AccreditationData,
		"escrow_id":            model.EscrowID,
		"kyc_at":               model.KycAt,
		"accreditation_at":     model.AccreditationAt,
		"created_at":           model.CreatedAt,
		"updated_at":           model.UpdatedAt,
		"wallet_id":            model.WalletID,
	}
}

func (model InvestmentProfile) Fields() []string {
	return []string{
		"ID",
		"UserID",
		"Type",
		"Data",
		"KycID",
		"KycStatus",
		"KycData",
		"AccreditationID",
		"AccreditationStatus",
		"AccreditationData",
		"EscrowID",
		"KycAt",
		"AccreditationAt",
		"CreatedAt",
		"UpdatedAt",
		"WalletID",
	}
}

func (model InvestmentProfile) Table() string {
	return "investment_profiles"
}

func (model InvestmentProfile) GetID() any {
	return model.ID
}

func (model *InvestmentProfile) SetID(id any) {
	model.ID = id.(int)
}
