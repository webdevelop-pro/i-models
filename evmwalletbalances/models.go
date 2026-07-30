package evmwalletbalances

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

// WalletBalance is the read-path cache for a managed wallet's token
// balance on a given chain.
type WalletBalance struct {
	ID                int                `db:"id" json:"id"`
	WalletAddress     string             `db:"wallet_address" json:"wallet_address"`
	Chain             string             `db:"chain" json:"chain"`
	TokenAddress      string             `db:"token_address" json:"token_address"`
	Balance           string             `db:"balance" json:"balance"`
	ObservedBlock     int64              `db:"observed_block" json:"observed_block"`
	ObservedBlockHash string             `db:"observed_block_hash" json:"observed_block_hash"`
	UpdatedAt         pgtype.Timestamptz `db:"updated_at" json:"updated_at"`

	db db.Repository `db:"-" json:"-"`
}

func (model WalletBalance) Fields() []string {
	return []string{
		"id", "wallet_address", "chain", "token_address", "balance",
		"observed_block", "observed_block_hash", "updated_at",
	}
}

func (model WalletBalance) Table() string {
	return TableName
}

func (model WalletBalance) ToJSON() map[string]any {
	res := map[string]any{}
	for _, field := range model.Fields() {
		res[field] = model.GetValueByTag(field)
	}
	return res
}

func (model WalletBalance) GetID() any {
	return model.ID
}

func (model *WalletBalance) SetID(id any) {
	model.ID = id.(int)
}

func (model *WalletBalance) SetDB(db db.Repository) {
	model.db = db
}

func (model WalletBalance) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "wallet_address":
		return model.WalletAddress
	case "chain":
		return model.Chain
	case "token_address":
		return model.TokenAddress
	case "balance":
		return model.Balance
	case "observed_block":
		return model.ObservedBlock
	case "observed_block_hash":
		return model.ObservedBlockHash
	case "updated_at":
		return model.UpdatedAt
	default:
		return nil
	}
}
