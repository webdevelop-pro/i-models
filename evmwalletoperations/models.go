package evmwalletoperations

import (
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

// WalletOperation is the canonical unified transaction record.
//
// TxHash is pgtype.Text so platform pre-insert rows can carry SQL NULL
// and avoid the partial unique index `WHERE tx_hash IS NOT NULL`.
// BlockTimestamp is pgtype.Timestamptz so it can be NULL until the
// webhook resolves it from the block.
type WalletOperation struct {
	ID                  int                `db:"id" json:"id"`
	UserID              int                `db:"user_id" json:"user_id"`
	ProfileID           int                `db:"profile_id" json:"profile_id"`
	InvestmentID        *int               `db:"investment_id" json:"investment_id"`
	TokenID             *int               `db:"token_id" json:"token_id"`
	WalletAddress       string             `db:"wallet_address" json:"wallet_address"`
	Chain               string             `db:"chain" json:"chain"`
	Source              string             `db:"source" json:"source"`
	Type                string             `db:"type" json:"type"`
	Status              string             `db:"status" json:"status"`
	Amount              string             `db:"amount" json:"amount"`
	AmountRaw           string             `db:"amount_raw" json:"amount_raw"`
	PriceUSD            string             `db:"price_usd" json:"price_usd"`
	AmountUSD           string             `db:"amount_usd" json:"amount_usd"`
	TokenTicker         string             `db:"token_ticker" json:"token_ticker"`
	TokenAddress        string             `db:"token_address" json:"token_address"`
	TokenName           string             `db:"token_name" json:"token_name"`
	TokenLogo           string             `db:"token_logo" json:"token_logo"`
	TokenDecimals       int                `db:"token_decimals" json:"token_decimals"`
	CounterpartyAddress string             `db:"counterparty_address" json:"counterparty_address"`
	TxHash              pgtype.Text        `db:"tx_hash" json:"tx_hash"`
	ExternalID          string             `db:"external_id" json:"external_id"`
	IdempotencyKey      string             `db:"idempotency_key" json:"idempotency_key"`
	ProviderEventID     string             `db:"provider_event_id" json:"provider_event_id"`
	FailureReason       string             `db:"failure_reason" json:"failure_reason"`
	BlockNumber         int64              `db:"block_number" json:"block_number"`
	BlockTimestamp      pgtype.Timestamptz `db:"block_timestamp" json:"block_timestamp"`
	ConfirmationCount   int                `db:"confirmation_count" json:"confirmation_count"`
	ConfirmationTarget  int                `db:"confirmation_target" json:"confirmation_target"`
	ReorgCount          int                `db:"reorg_count" json:"reorg_count"`
	LastSeenBlock       int64              `db:"last_seen_block" json:"last_seen_block"`
	RemovedAt           pgtype.Timestamptz `db:"removed_at" json:"removed_at"`
	CreatedAt           pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt           pgtype.Timestamptz `db:"updated_at" json:"updated_at"`

	db db.Repository `db:"-" json:"-"`
}

func (model WalletOperation) Fields() []string {
	return models.DefaultFields(&model)
}

func (model WalletOperation) Table() string {
	return TableName
}

func (model WalletOperation) ToJSON() map[string]any {
	res := map[string]any{}
	for _, field := range model.Fields() {
		res[field] = model.GetValueByTag(field)
	}
	return res
}

func (model WalletOperation) GetID() any {
	return model.ID
}

func (model *WalletOperation) SetID(id any) {
	model.ID = id.(int)
}

func (model *WalletOperation) SetDB(db db.Repository) {
	model.db = db
}

func (model WalletOperation) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "profile_id":
		return model.ProfileID
	case "investment_id":
		return model.InvestmentID
	case "token_id":
		return model.TokenID
	case "wallet_address":
		return model.WalletAddress
	case "chain":
		return model.Chain
	case "source":
		return model.Source
	case "type":
		return model.Type
	case "status":
		return model.Status
	case "amount":
		return model.Amount
	case "amount_raw":
		return model.AmountRaw
	case "price_usd":
		return model.PriceUSD
	case "amount_usd":
		return model.AmountUSD
	case "token_ticker":
		return model.TokenTicker
	case "token_address":
		return model.TokenAddress
	case "token_name":
		return model.TokenName
	case "token_logo":
		return model.TokenLogo
	case "token_decimals":
		return model.TokenDecimals
	case "counterparty_address":
		return model.CounterpartyAddress
	case "tx_hash":
		return model.TxHash
	case "external_id":
		return model.ExternalID
	case "idempotency_key":
		return model.IdempotencyKey
	case "provider_event_id":
		return model.ProviderEventID
	case "failure_reason":
		return model.FailureReason
	case "block_number":
		return model.BlockNumber
	case "block_timestamp":
		return model.BlockTimestamp
	case "confirmation_count":
		return model.ConfirmationCount
	case "confirmation_target":
		return model.ConfirmationTarget
	case "reorg_count":
		return model.ReorgCount
	case "last_seen_block":
		return model.LastSeenBlock
	case "removed_at":
		return model.RemovedAt
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	default:
		return nil
	}
}
