package evmwalletoperationeffects

type EffectKindT string
type EffectDirectionT string

const (
	EffectKindNative        EffectKindT = "native"
	EffectKindERC20Transfer EffectKindT = "erc20_transfer"
	EffectKindContractEvent EffectKindT = "contract_event"
)

const (
	EffectDirectionDeposit    EffectDirectionT = "deposit"
	EffectDirectionWithdrawal EffectDirectionT = "withdrawal"
)

func AllEffectKindT() []EffectKindT {
	return []EffectKindT{EffectKindNative, EffectKindERC20Transfer, EffectKindContractEvent}
}

func AllEffectDirectionT() []EffectDirectionT {
	return []EffectDirectionT{EffectDirectionDeposit, EffectDirectionWithdrawal}
}
