package wallets

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

const Table = "wallet_wallets"

// WalletWallet is an object representing the database table.
type WalletWallet struct {
	ID                      int                `db:"-" json:"id" yaml:"id"`
	EntityCustomerID        *string            `db:"-" json:"entity_customer_id,omitempty" yaml:"entity_customer_id,omitempty"`
	EntityCustomerBalanceID *string            `db:"-" json:"entity_customer_balance_id,omitempty" yaml:"entity_customer_balance_id,omitempty"`
	Balance                 float64            `db:"-" json:"balance" yaml:"balance"`
	IncBalance              float64            `db:"-" json:"inc_balance" yaml:"inc_balance"`
	OutBalance              float64            `db:"-" json:"out_balance" yaml:"out_balance"`
	Status                  WalletStatusT      `db:"-" json:"status" yaml:"status"`
	CreatedAt               pgtype.Timestamptz `db:"-" json:"created_at" yaml:"created_at"`
	UpdatedAt               pgtype.Timestamptz `db:"-" json:"updated_at" yaml:"updated_at"`
}

func (model WalletWallet) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "EntityCustomerID":
		return model.EntityCustomerID
	case "EntityCustomerBalanceID":
		return model.EntityCustomerBalanceID
	case "Balance":
		return model.Balance
	case "IncBalance":
		return model.IncBalance
	case "OutBalance":
		return model.OutBalance
	case "Status":
		return model.Status
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model WalletWallet) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetField(key)
	}
	return res
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

func (model *WalletWallet) SetID(id any) {
	model.ID = id.(int)
}
