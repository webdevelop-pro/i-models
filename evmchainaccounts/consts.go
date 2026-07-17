package evmchainaccounts

type ChainT string
type AccountModeT string
type StatusT string

const (
	ChainEthereum        ChainT = "ethereum"
	ChainPolygon         ChainT = "polygon"
	ChainOptimism        ChainT = "optimism"
	ChainBase            ChainT = "base"
	ChainEthereumSepolia ChainT = "ethereum-sepolia"
)

const (
	AccountMode7702 AccountModeT = "7702"
	AccountModeSCA  AccountModeT = "sca"
)

const (
	StatusPending   StatusT = "pending"
	StatusVerified  StatusT = "verified"
	StatusFailed    StatusT = "failed"
	StatusSuspended StatusT = "suspended"
)

func AllChainT() []ChainT {
	return []ChainT{ChainEthereum, ChainPolygon, ChainOptimism, ChainBase, ChainEthereumSepolia}
}

func AllAccountModeT() []AccountModeT {
	return []AccountModeT{AccountMode7702, AccountModeSCA}
}

func AllStatusT() []StatusT {
	return []StatusT{StatusPending, StatusVerified, StatusFailed, StatusSuspended}
}

const (
	AppLabel  = "evm"
	ModelName = "chain_accounts"
	TableName = "evm_chain_accounts"

	pkgName = "models/evmchainaccounts"
)
