package evmwalletbalances

import (
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// WalletBalance is the read-path cache for a managed wallet's token
// balance on a given chain.
type WalletBalance struct {
	ID            int                `db:"id" json:"id"`
	WalletAddress string             `db:"wallet_address" json:"wallet_address"`
	Chain         string             `db:"chain" json:"chain"`
	TokenAddress  string             `db:"token_address" json:"token_address"`
	Balance       string             `db:"balance" json:"balance"`
	UpdatedAt     pgtype.Timestamptz `db:"updated_at" json:"updated_at"`

	db db.Repository `db:"-" json:"-"`
}

func (model WalletBalance) Fields() []string {
	return models.DefaultFields(&model)
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
	case "updated_at":
		return model.UpdatedAt
	default:
		return nil
	}
}
