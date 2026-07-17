package evmchainaccounts

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

// EvmChainAccount is an object representing the database table.
type EvmChainAccount struct {
	ID        int  `db:"id" json:"id" yaml:"id"`
	UserID    *int `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	WalletID  int  `db:"wallet_id" json:"wallet_id" yaml:"wallet_id"`
	ProfileID *int `db:"profile_id" json:"profile_id,omitempty" yaml:"profile_id,omitempty"`

	Chain              ChainT       `db:"chain" json:"chain" yaml:"chain"`
	AccountMode        AccountModeT `db:"account_mode" json:"account_mode" yaml:"account_mode"`
	Address            string       `db:"address" json:"address" yaml:"address"`
	SignerAddress      string       `db:"signer_address" json:"signer_address" yaml:"signer_address"`
	AlchemyOwnerID     string       `db:"alchemy_owner_id" json:"-" yaml:"-"`
	AlchemyAccountID   string       `db:"alchemy_account_id" json:"-" yaml:"-"`
	MonitoringID       string       `db:"monitoring_id" json:"monitoring_id" yaml:"monitoring_id"`
	Status             StatusT      `db:"status" json:"status" yaml:"status"`
	ProviderName       string       `db:"provider_name" json:"provider_name" yaml:"provider_name"`
	TurnkeyAccountID   string       `db:"turnkey_account_id" json:"-" yaml:"-"`
	MonitoringProvider string       `db:"monitoring_provider" json:"-" yaml:"-"`

	ActivatedAt pgtype.Timestamptz `db:"activated_at" json:"activated_at" yaml:"activated_at"`
	CreatedAt   pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func NewEvmChainAccount(db db.Repository) *EvmChainAccount {
	return &EvmChainAccount{
		db:  db,
		fns: map[string]any{},
	}
}

func (model EvmChainAccount) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "WalletID":
		return model.WalletID
	case "ProfileID":
		return model.ProfileID
	case "Chain":
		return model.Chain
	case "AccountMode":
		return model.AccountMode
	case "Address":
		return model.Address
	case "SignerAddress":
		return model.SignerAddress
	case "MonitoringID":
		return model.MonitoringID
	case "Status":
		return model.Status
	case "ActivatedAt":
		return model.ActivatedAt
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model EvmChainAccount) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "wallet_id":
		return model.WalletID
	case "profile_id":
		return model.ProfileID
	case "chain":
		return model.Chain
	case "account_mode":
		return model.AccountMode
	case "address":
		return model.Address
	case "signer_address":
		return model.SignerAddress
	case "monitoring_id":
		return model.MonitoringID
	case "status":
		return model.Status
	case "activated_at":
		return model.ActivatedAt
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	}
	return nil
}

func (model EvmChainAccount) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetValueByTag(key)
	}
	return res
}

func (model EvmChainAccount) Fields() []string {
	return []string{
		"id", "user_id", "wallet_id", "profile_id", "chain", "account_mode",
		"address", "signer_address", "monitoring_id", "status", "activated_at",
		"created_at", "updated_at",
	}
}

func (model EvmChainAccount) Table() string {
	return TableName
}

func (model EvmChainAccount) GetID() any {
	return model.ID
}

func (model *EvmChainAccount) SetID(id any) {
	model.ID = id.(int)
}

func (model *EvmChainAccount) SetDB(db db.Repository) {
	model.db = db
}
