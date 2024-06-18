package offers

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// ToDo
// Refactor, create standalone field for the escrow data
type Data struct {
	WireTo                string `json:"wire_to"`
	SwiftID               string `json:"swift_id"`
	Custodian             string `json:"custodian"`
	AccountNumber         string `json:"account_number"`
	RoutingNumber         string `json:"routing_number"`
	NCStampingText        string `json:"nc_stamping_text"`
	NCEscrowAccountNumber string `json:"nc_escrow_account_number"`
	NCEscrowRoutingNumber string `json:"nc_escrow_routing_number"`
}

// ToDo
// Add fields
type Offer struct {
	ID            int          `json:"id"`
	WalletID      *int         `json:"wallet_id"`
	UserID        int          `json:"user_id"`
	Name          string       `json:"name"`
	MinInvestment float64      `json:"min_investment"`
	Description   string       `json:"description"`
	Valuation     float64      `json:"valuation"`
	TotalShares   int          `json:"total_shares"`
	PricePerShare float64      `json:"price_per_share"`
	EntityID      string       `json:"entity_id"`
	SecurityType  SecurityType `json:"security_type"`
	Data          Data         `json:"data"`
	Status        Type         `json:"status"`
	RegType       RegType      `json:"reg_type"`

	StartAt   pgtype.Timestamptz `json:"start_at"`
	CloseAt   pgtype.Timestamptz `json:"close_at"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
