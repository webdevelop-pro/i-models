package fundingsources

import (
	"github.com/webdevelop-pro/i-models/pgtype"
	"github.com/webdevelop-pro/i-models/wallets"
)

// WalletFundingSource is an object representing the database table.
type WalletFundingSource struct {
	ID        int                   `json:"id" yaml:"id"`
	WalletID  int                   `json:"wallet_id,omitempty" yaml:"wallet_id,omitempty"`
	EntityID  string                `json:"entity_id" yaml:"entity_id"`
	Type      FoundingSourceT       `json:"type" yaml:"type"`
	BankName  string                `json:"bank_name" yaml:"bank_name"`
	Status    wallets.WalletStatusT `json:"status" yaml:"status"`
	Name      string                `json:"name" yaml:"name"`
	CreatedAt pgtype.Timestamptz    `json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz    `json:"updated_at" yaml:"updated_at"`
}

func (model WalletFundingSource) ToJSON() map[string]any {
	return map[string]any{
		"id":         model.ID,
		"wallet_id":  model.WalletID,
		"entity_id":  model.EntityID,
		"type":       model.Type,
		"bank_name":  model.BankName,
		"status":     model.Status,
		"name":       model.Name,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
	}
}

func (model WalletFundingSource) Fields() []string {
	return []string{
		"ID",
		"WalletID",
		"EntityID",
		"Type",
		"BankName",
		"Status",
		"Name",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model WalletFundingSource) Table() string {
	return "wallet_funding_source"
}

func (model WalletFundingSource) GetID() any {
	return model.ID
}
