package investments

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// InvestmentInvestment is an object representing the database table.
type InvestmentInvestment struct {
	ID                int                `json:"id" yaml:"id"`
	UserID            int                `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	OfferID           int                `json:"offer_id,omitempty" yaml:"offer_id,omitempty"`
	ProfileID         int                `json:"profile_id,omitempty" yaml:"profile_id,omitempty"`
	Amount            float64            `json:"amount" yaml:"amount"`
	PricePerShare     float64            `json:"price_per_share" yaml:"price_per_share"`
	NumberOfShares    int                `json:"number_of_shares" yaml:"number_of_shares"`
	PaymentType       PaymentT           `json:"payment_type" yaml:"payment_type"`
	EscrowType        EscrowT            `json:"escrow_type" yaml:"escrow_type"`
	FundingType       FundingT           `json:"funding_type" yaml:"funding_type"`
	FundingStatus     FundingS           `json:"funding_status" yaml:"funding_status"`
	Status            InvestmentT        `json:"status" yaml:"status"`
	PrevStatus        InvestmentT        `json:"prev_status" yaml:"prev_status"`
	Step              InvestmentStepT    `json:"step" yaml:"step"`
	Commission        float64            `json:"commission" yaml:"commission"`
	CancelationReason string             `json:"cancelation_reason" yaml:"cancelation_reason"`
	EntityID          *string            `json:"entity_id,omitempty" yaml:"entity_id,omitempty"`
	EscrowData        any                `json:"escrow_data" yaml:"escrow_data"`
	SignatureData     any                `json:"signature_data" yaml:"signature_data"`
	PaymentData       any                `json:"payment_data" yaml:"payment_data"`
	Log               any                `json:"log" yaml:"log"`
	CanceledAt        pgtype.Timestamptz `json:"canceled_at" yaml:"canceled_at"`
	SubmitedAt        pgtype.Timestamptz `json:"submited_at" yaml:"submited_at"`
	ClosedAt          pgtype.Timestamptz `json:"closed_at" yaml:"closed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
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
		"escrow_data":        model.EscrowData,
		"signature_data":     model.SignatureData,
		"payment_data":       model.PaymentData,
		"log":                model.Log,
		"canceled_at":        model.CanceledAt,
		"submited_at":        model.SubmitedAt,
		"closed_at":          model.ClosedAt,
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
	}
}

func (model InvestmentInvestment) Fields() []string {
	return []string{
		"ID",
		"UserID",
		"OfferID",
		"ProfileID",
		"Amount",
		"PricePerShare",
		"NumberOfShares",
		"PaymentType",
		"EscrowType",
		"FundingType",
		"FundingStatus",
		"Status",
		"PrevStatus",
		"Step",
		"Commission",
		"CancelationReason",
		"EntityID",
		"EscrowData",
		"SignatureData",
		"PaymentData",
		"Log",
		"CanceledAt",
		"SubmitedAt",
		"ClosedAt",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model InvestmentInvestment) Table() string {
	return "investment_investments"
}

func (model InvestmentInvestment) GetID() any {
	return model.ID
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
