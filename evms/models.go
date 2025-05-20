package fundingsources

import (
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// Wallet is an object representing the database table.
type Wallet struct {
	ID     int `db:"id" json:"id" yaml:"id"`
	UserID int `db:"user_id" json:"user_id" yaml:"user_id"`

	PublicKey  string `db:"public_key" json:"public_key" yaml:"public_key"`
	PrivateKey string `db:"private_key" json:"private_key" yaml:"private_key"`

	Balance    float64            `db:"-" json:"balance" yaml:"balance"`
	IncBalance float64            `db:"-" json:"inc_balance" yaml:"inc_balance"`
	OutBalance float64            `db:"-" json:"out_balance" yaml:"out_balance"`
	Status     WalletStatusT      `db:"-" json:"status" yaml:"status"`
	CreatedAt  pgtype.Timestamptz `db:"-" json:"created_at" yaml:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"-" json:"updated_at" yaml:"updated_at"`

	ObjectID      string `db:"-" json:"object_id" yaml:"object_id"`
	ContentTypeID int    `db:"-" json:"content_type_id" yaml:"content_type_id"`
}

func (model Wallet) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "PublicKey":
		return model.PublicKey
	case "PrivateKey":
		return model.PrivateKey
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

func (model Wallet) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetField(key)
	}
	return res
}

func (model Wallet) Fields() []string {
	return models.DefaultFields(&model)
}

func (model Wallet) Table() string {
	return TableName
}

func (model Wallet) GetID() any {
	return model.ID
}

func (model *Wallet) SetID(id any) {
	model.ID = id.(int)
}
