package transactions

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

// WalletTransaction is an object representing the database table.
type WalletTransaction struct {
	ID             int                 `json:"id" yaml:"id"`
	SourceWalletID int                 `json:"source_wallet_id,omitempty" yaml:"source_wallet_id,omitempty"`
	DestWalletID   int                 `json:"dest_wallet_id,omitempty" yaml:"dest_wallet_id,omitempty"`
	InvestmentID   int                 `json:"investment_id,omitempty" yaml:"investment_id,omitempty"`
	EntityID       *string             `json:"entity_id,omitempty" yaml:"entity_id,omitempty"`
	Type           TransactionsTypeT   `json:"type" yaml:"type"`
	Amount         float64             `json:"amount" yaml:"amount"`
	Status         TransactionsStatusT `json:"status" yaml:"status"`
	CreatedAt      pgtype.Timestamptz  `json:"created_at" yaml:"created_at"`
	UpdatedAt      pgtype.Timestamptz  `json:"updated_at" yaml:"updated_at"`
}

func (model WalletTransaction) ToJSON() map[string]any {
	return map[string]any{
		"id":               model.ID,
		"source_wallet_id": model.SourceWalletID,
		"dest_wallet_id":   model.DestWalletID,
		"investment_id":    model.InvestmentID,
		"entity_id":        model.EntityID,
		"type":             model.Type,
		"amount":           model.Amount,
		"status":           model.Status,
		"created_at":       model.CreatedAt,
		"updated_at":       model.UpdatedAt,
	}
}

func (model WalletTransaction) Fields() []string {
	return []string{
		"ID",
		"SourceWalletID",
		"DestWalletID",
		"InvestmentID",
		"EntityID",
		"Type",
		"Amount",
		"Status",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model WalletTransaction) Table() string {
	return "wallet_transactions"
}

func (model WalletTransaction) GetID() any {
	return model.ID
}
