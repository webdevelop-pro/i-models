package wallets

import "github.com/jackc/pgtype"

type Wallet struct {
	ID                      int    `json:"id"`
	EntityCustomerID        string `json:"entity_customer_id"`
	EntityCustomerBalanceID string `json:"entity_customer_balance_id"`
	Balance                 string `json:"balance"`
	Status                  Status `json:"status"`

	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
