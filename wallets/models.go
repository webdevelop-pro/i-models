package wallets

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// WalletWallet is an object representing the database table.
type WalletWallet struct {
	ID                      int           `json:"id" yaml:"id"`
	EntityCustomerID        *string       `json:"entity_customer_id,omitempty" yaml:"entity_customer_id,omitempty"`
	EntityCustomerBalanceID *string       `json:"entity_customer_balance_id,omitempty" yaml:"entity_customer_balance_id,omitempty"`
	Balance                 float64       `json:"balance" yaml:"balance"`
	IncBalance              float64       `json:"inc_balance" yaml:"inc_balance"`
	OutBalance              float64       `json:"out_balance" yaml:"out_balance"`
	Status                  WalletStatusT `json:"status" yaml:"status"`
	// Log                     any                `json:"log" yaml:"log"`
	CreatedAt pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
}

func (model WalletWallet) ToJSON() map[string]any {
	return map[string]any{
		"id":                         model.ID,
		"entity_customer_id":         model.EntityCustomerID,
		"entity_customer_balance_id": model.EntityCustomerBalanceID,
		"balance":                    model.Balance,
		"inc_balance":                model.IncBalance,
		"out_balance":                model.OutBalance,
		"status":                     model.Status,
		"created_at":                 model.CreatedAt,
		"updated_at":                 model.UpdatedAt,
		// "log":                        model.Log,

	}
}

func (model WalletWallet) Fields() []string {
	return []string{
		"ID",
		"EntityCustomerID",
		"EntityCustomerBalanceID",
		"Balance",
		"IncBalance",
		"OutBalance",
		"Status",
		// "Log",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model WalletWallet) Table() string {
	return "wallet_wallets"
}

func (model WalletWallet) GetID() any {
	return model.ID
}
