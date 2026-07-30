package evmallowancesequences

import "github.com/global-torque/go-common/orm/v2/pgtype"

// AllowanceSequence durably coordinates the approve(0), approve(amount),
// pull, and cleanup operations for one exact ERC-20 allowance tuple.
type AllowanceSequence struct {
	ID                     int64              `db:"id" json:"id"`
	Purpose                PurposeT           `db:"purpose" json:"purpose"`
	CustodyChainAccountID  int                `db:"custody_chain_account_id" json:"custody_chain_account_id"`
	Chain                  string             `db:"chain" json:"chain"`
	TokenAddress           string             `db:"token_address" json:"token_address"`
	SpenderContractID      int                `db:"spender_contract_id" json:"spender_contract_id"`
	InvestmentID           *int               `db:"investment_id" json:"investment_id,omitempty"`
	InvestmentRedemptionID *int64             `db:"investment_redemption_id" json:"investment_redemption_id,omitempty"`
	IdempotencyKey         string             `db:"idempotency_key" json:"idempotency_key"`
	ApprovalOperationID    *int               `db:"approval_operation_id" json:"approval_operation_id,omitempty"`
	PullOperationID        *int               `db:"pull_operation_id" json:"pull_operation_id,omitempty"`
	CleanupOperationID     *int               `db:"cleanup_operation_id" json:"cleanup_operation_id,omitempty"`
	AcquiredAt             pgtype.Timestamptz `db:"acquired_at" json:"acquired_at"`
	CompletedAt            pgtype.Timestamptz `db:"completed_at" json:"completed_at,omitempty"`
	CancelledAt            pgtype.Timestamptz `db:"cancelled_at" json:"cancelled_at,omitempty"`
	LastError              string             `db:"last_error" json:"last_error"`
	CreatedAt              pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

func (model AllowanceSequence) ToJSON() map[string]any {
	return map[string]any{
		"id":                       model.ID,
		"purpose":                  model.Purpose,
		"custody_chain_account_id": model.CustodyChainAccountID,
		"chain":                    model.Chain,
		"token_address":            model.TokenAddress,
		"spender_contract_id":      model.SpenderContractID,
		"investment_id":            model.InvestmentID,
		"investment_redemption_id": model.InvestmentRedemptionID,
		"idempotency_key":          model.IdempotencyKey,
		"approval_operation_id":    model.ApprovalOperationID,
		"pull_operation_id":        model.PullOperationID,
		"cleanup_operation_id":     model.CleanupOperationID,
		"acquired_at":              model.AcquiredAt,
		"completed_at":             model.CompletedAt,
		"cancelled_at":             model.CancelledAt,
		"last_error":               model.LastError,
		"created_at":               model.CreatedAt,
		"updated_at":               model.UpdatedAt,
	}
}

func (model AllowanceSequence) Fields() []string {
	return []string{
		"id", "purpose", "custody_chain_account_id", "chain", "token_address",
		"spender_contract_id", "investment_id", "investment_redemption_id",
		"idempotency_key", "approval_operation_id", "pull_operation_id",
		"cleanup_operation_id",
		"acquired_at", "completed_at", "cancelled_at",
		"last_error", "created_at", "updated_at",
	}
}

func (model AllowanceSequence) Table() string { return "evm_erc20_allowance_sequences" }
func (model AllowanceSequence) GetID() any    { return model.ID }
func (model *AllowanceSequence) SetID(id any) { model.ID = id.(int64) }
