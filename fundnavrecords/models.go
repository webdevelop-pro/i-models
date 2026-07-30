package fundnavrecords

import "github.com/global-torque/go-common/orm/v2/pgtype"

// FundNAVRecord stores the complete, reproducible valuation snapshot used to
// price ERC-7540 deposits and redemptions. Numeric values stay as decimal
// strings so uint256-scale values do not pass through floating point.
type FundNAVRecord struct {
	ID                           int64              `db:"id" json:"id"`
	OfferID                      int                `db:"offer_id" json:"offer_id"`
	NAV                          string             `db:"nav" json:"nav"`
	RecordedAt                   pgtype.Timestamptz `db:"recorded_at" json:"recorded_at"`
	RecordedByUserID             *int               `db:"recorded_by_user_id" json:"recorded_by_user_id,omitempty"`
	RecordedByName               string             `db:"recorded_by_name" json:"recorded_by_name"`
	Source                       string             `db:"source" json:"source"`
	CreatedAt                    pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt                    pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	Version                      *int64             `db:"version" json:"version,omitempty"`
	ExternalStrategyAssetsRaw    string             `db:"external_strategy_assets_raw" json:"external_strategy_assets_raw"`
	OfferWalletUSDCRaw           *string            `db:"offer_wallet_usdc_raw" json:"offer_wallet_usdc_raw,omitempty"`
	CustodyPrincipalReserveRaw   *string            `db:"custody_principal_reserve_raw" json:"custody_principal_reserve_raw,omitempty"`
	VaultLiquidAssetsRaw         *string            `db:"vault_liquid_assets_raw" json:"vault_liquid_assets_raw,omitempty"`
	VaultTotalSupplyRaw          *string            `db:"vault_total_supply_raw" json:"vault_total_supply_raw,omitempty"`
	ClaimableRedemptionAssetsRaw *string            `db:"claimable_redemption_assets_raw" json:"claimable_redemption_assets_raw,omitempty"`
	ClaimableRedemptionSharesRaw *string            `db:"claimable_redemption_shares_raw" json:"claimable_redemption_shares_raw,omitempty"`
	StrategyWalletManagedUSDCRaw *string            `db:"strategy_wallet_managed_usdc_raw" json:"strategy_wallet_managed_usdc_raw,omitempty"`
	StrategyAssetsRaw            *string            `db:"strategy_assets_raw" json:"strategy_assets_raw,omitempty"`
	NAVAssetsRaw                 *string            `db:"nav_assets_raw" json:"nav_assets_raw,omitempty"`
	NAVShareSupplyRaw            *string            `db:"nav_share_supply_raw" json:"nav_share_supply_raw,omitempty"`
	NAVUSDCRaw                   *string            `db:"nav_usdc_raw" json:"nav_usdc_raw,omitempty"`
	ValuationBlockNumber         *int64             `db:"valuation_block_number" json:"valuation_block_number,omitempty"`
	ValuationBlockHash           *string            `db:"valuation_block_hash" json:"valuation_block_hash,omitempty"`
	ValuationAsOf                pgtype.Timestamptz `db:"valuation_as_of" json:"valuation_as_of,omitempty"`
	FinalizedAt                  pgtype.Timestamptz `db:"finalized_at" json:"finalized_at,omitempty"`
}

func (model FundNAVRecord) ToJSON() map[string]any {
	return map[string]any{
		"id":                               model.ID,
		"offer_id":                         model.OfferID,
		"nav":                              model.NAV,
		"recorded_at":                      model.RecordedAt,
		"recorded_by_user_id":              model.RecordedByUserID,
		"recorded_by_name":                 model.RecordedByName,
		"source":                           model.Source,
		"created_at":                       model.CreatedAt,
		"updated_at":                       model.UpdatedAt,
		"version":                          model.Version,
		"external_strategy_assets_raw":     model.ExternalStrategyAssetsRaw,
		"offer_wallet_usdc_raw":            model.OfferWalletUSDCRaw,
		"custody_principal_reserve_raw":    model.CustodyPrincipalReserveRaw,
		"vault_liquid_assets_raw":          model.VaultLiquidAssetsRaw,
		"vault_total_supply_raw":           model.VaultTotalSupplyRaw,
		"claimable_redemption_assets_raw":  model.ClaimableRedemptionAssetsRaw,
		"claimable_redemption_shares_raw":  model.ClaimableRedemptionSharesRaw,
		"strategy_wallet_managed_usdc_raw": model.StrategyWalletManagedUSDCRaw,
		"strategy_assets_raw":              model.StrategyAssetsRaw,
		"nav_assets_raw":                   model.NAVAssetsRaw,
		"nav_share_supply_raw":             model.NAVShareSupplyRaw,
		"nav_usdc_raw":                     model.NAVUSDCRaw,
		"valuation_block_number":           model.ValuationBlockNumber,
		"valuation_block_hash":             model.ValuationBlockHash,
		"valuation_as_of":                  model.ValuationAsOf,
		"finalized_at":                     model.FinalizedAt,
	}
}

func (model FundNAVRecord) Fields() []string {
	return []string{
		"id", "offer_id", "nav", "recorded_at", "recorded_by_user_id",
		"recorded_by_name", "source", "created_at", "updated_at", "version",
		"external_strategy_assets_raw", "offer_wallet_usdc_raw",
		"custody_principal_reserve_raw", "vault_liquid_assets_raw",
		"vault_total_supply_raw", "claimable_redemption_assets_raw",
		"claimable_redemption_shares_raw", "strategy_wallet_managed_usdc_raw",
		"strategy_assets_raw", "nav_assets_raw", "nav_share_supply_raw",
		"nav_usdc_raw", "valuation_block_number", "valuation_block_hash",
		"valuation_as_of", "finalized_at",
	}
}

func (model FundNAVRecord) Table() string { return "fund_nav_records" }
func (model FundNAVRecord) GetID() any    { return model.ID }
func (model *FundNAVRecord) SetID(id any) { model.ID = id.(int64) }
