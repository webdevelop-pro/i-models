package investments

import "github.com/jackc/pgtype"

// ToDo
// Add all other fields
type Investment struct {
	ID             int           `json:"id"`
	UserID         int           `json:"user_id"`
	ProfileID      int           `json:"profile_id"`
	NumberOfShares int           `json:"number_of_shares"`
	Amount         float64       `json:"amount"`
	PricePerShare  float64       `json:"price_per_share"`
	Commission     float64       `json:"commission"`
	PaymentType    PaymentType   `json:"payment_type"`
	EscrowType     EscrowType    `json:"escrow_type"`
	FundingType    FundingType   `json:"funding_type"`
	FundingStatus  FundingStatus `json:"funding_status"`
	Status         Status        `json:"status"`
	PrevStatus     Status        `json:"prev_status"`
	Step           string        `json:"step"`
	EntityID       string        `json:"entity_id"`

	CanceledAt pgtype.Timestamptz `json:"canceled_at"`
	SubmitedAt pgtype.Timestamptz `json:"submited_at"`
	ClosedAt   pgtype.Timestamptz `json:"closed_at"`
	CreatedAt  pgtype.Timestamptz `json:"created_at,omitempty"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at,omitempty"`
}

// IsFundingTypeWire temporary fix for https://github.com/Joker/jade/issues/59
func (inv *Investment) IsFundingTypeWire() bool {
	return inv.FundingType == "wire"
}
