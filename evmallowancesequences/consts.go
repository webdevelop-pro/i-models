package evmallowancesequences

type PurposeT string

const (
	PurposeDepositRequest PurposeT = "deposit_request"
	PurposeStrategyReturn PurposeT = "strategy_return"
)

func AllPurposeT() []PurposeT {
	return []PurposeT{
		PurposeDepositRequest,
		PurposeStrategyReturn,
	}
}
