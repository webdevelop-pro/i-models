package evmwalletoperations

type OperationTypeT string
type OperationStatusT string
type OperationSourceT string

const (
	OperationTypeDeposit    OperationTypeT = "deposit"
	OperationTypeWithdrawal OperationTypeT = "withdrawal"
	OperationTypeInvestment OperationTypeT = "investment"
)

const (
	OperationStatusCreated   OperationStatusT = "created"
	OperationStatusSubmitted OperationStatusT = "submitted"
	OperationStatusConfirmed OperationStatusT = "confirmed"
	OperationStatusFailed    OperationStatusT = "failed"
)

const (
	OperationSourceWebhook  OperationSourceT = "webhook"
	OperationSourcePlatform OperationSourceT = "platform"
)

func AllOperationTypeT() []OperationTypeT {
	return []OperationTypeT{OperationTypeDeposit, OperationTypeWithdrawal, OperationTypeInvestment}
}

func AllOperationStatusT() []OperationStatusT {
	return []OperationStatusT{
		OperationStatusCreated,
		OperationStatusSubmitted,
		OperationStatusConfirmed,
		OperationStatusFailed,
	}
}

func AllOperationSourceT() []OperationSourceT {
	return []OperationSourceT{OperationSourceWebhook, OperationSourcePlatform}
}

const (
	AppLabel  = "evm"
	ModelName = "wallet_operations"
	TableName = "evm_wallet_operations"

	pkgName = "models/evmwalletoperations"
)
