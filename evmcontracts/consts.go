package evmcontracts

import "github.com/pkg/errors"

type StatusT string

type DeploymentLegT string

// Enum values for StatusT
const (
	StatusTCreated  StatusT = "created"
	StatusTPending  StatusT = "pending"
	StatusTDeployed StatusT = "deployed"
	StatusTVerified StatusT = "verified"
	StatusTError    StatusT = "error"

	AppLabel                              = "evm"
	ModelName                             = "contract"
	TableName                             = "evm_contracts"
	DeploymentOperationsTableName         = "evm_contract_deployment_operations"
	DeploymentTransactionHistoryTableName = "evm_contract_deployment_transaction_history"
	SignerNonceReservationsTableName      = "evm_signer_nonce_reservations"

	pkgName = "models/emvcontracts"

	DeploymentLegTToken             DeploymentLegT = "token"
	DeploymentLegTTreasury          DeploymentLegT = "treasury"
	DeploymentLegTTreasuryAllowlist DeploymentLegT = "treasury_allowlist"
	DeploymentLegTTreasuryFunding   DeploymentLegT = "treasury_funding"
	DeploymentLegTVault             DeploymentLegT = "vault"
	DeploymentLegTUSDCFunding       DeploymentLegT = "usdc_funding"
	DeploymentLegTExchangePayout    DeploymentLegT = "exchange_payout"
)

func AllStatusT() []StatusT {
	return []StatusT{
		StatusTCreated,
		StatusTPending,
		StatusTDeployed,
		StatusTVerified,
		StatusTError,
	}
}

func (e StatusT) IsValid() error {
	switch e {
	case StatusTCreated, StatusTPending, StatusTDeployed, StatusTVerified, StatusTError:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e StatusT) String() string {
	return string(e)
}
