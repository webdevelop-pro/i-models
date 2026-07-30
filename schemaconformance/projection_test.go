package schemaconformance_test

import (
	"context"
	"os"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/webdevelop-pro/go-common/logger"

	"github.com/webdevelop-pro/i-models/fundnavrecords"
	"github.com/webdevelop-pro/i-models/investments"
	"github.com/webdevelop-pro/i-models/offers"
)

type transactionRepository struct {
	pgx.Tx
}

func (transactionRepository) Lg() logger.Logger { return logger.Logger{} }

type offerProjection struct {
	ID                    int                `db:"id"`
	Status                string             `db:"status"`
	Data                  map[string]any     `db:"data"`
	InvestorEligibility   []string           `db:"investor_eligibility"`
	MarketedJurisdictions []string           `db:"marketed_jurisdictions"`
	RiskDisclosures       string             `db:"risk_disclosures"`
	StartAt               pgtype.Timestamptz `db:"start_at"`
	TotalShares           string             `db:"total_shares"`
	PricePerShare         string             `db:"price_per_share"`
	DataRoomGroupID       *int               `db:"data_room_group_id"`
	FundStructure         string             `db:"fund_structure"`
}

func TestOfferProjectionAgainstPostgreSQL(t *testing.T) {
	if os.Getenv("I_MODELS_SCHEMA_TEST") != "1" {
		t.Skip("set I_MODELS_SCHEMA_TEST=1 to run PostgreSQL projection parity")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, schemaDatabaseURL())
	if err != nil {
		t.Fatalf("connect to projection database: %v", err)
	}
	t.Cleanup(func() {
		if err := conn.Close(ctx); err != nil {
			t.Errorf("close projection database: %v", err)
		}
	})

	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatalf("begin projection transaction: %v", err)
	}
	t.Cleanup(func() {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			t.Errorf("rollback projection transaction: %v", err)
		}
	})
	repo := transactionRepository{Tx: tx}

	var userID int
	err = tx.QueryRow(ctx, `
		INSERT INTO user_users(id, password, email, first_name, last_name, identity_id)
		VALUES (-700001, '', 'projection@example.test', 'Projection', 'Test', 'projection-schema-test')
		RETURNING id
	`).Scan(&userID)
	if err != nil {
		t.Fatalf("insert projection user: %v", err)
	}

	var offerID int
	err = tx.QueryRow(ctx, `
		INSERT INTO offer_offers(
			user_id, name, slug, min_investment, status, data, investor_eligibility,
			marketed_jurisdictions, risk_disclosures, start_at, total_shares, price_per_share
		) VALUES (
			$1, $2, $3, $4, $5, $6::jsonb, $7::jsonb, $8::jsonb, $9, NULL, $10, $11
		) RETURNING id
	`,
		userID,
		"Projection offer",
		"projection-offer",
		100,
		"legal_review",
		`{"nc_stamping_text":"Stamped \"quoted\" ✓"}`,
		`["accredited","qualified_investor"]`,
		`["US","EU"]`,
		"Risk disclosure with Unicode ✓",
		125.25,
		2.5,
	).Scan(&offerID)
	if err != nil {
		t.Fatalf("insert projection offer: %v", err)
	}

	projection, err := orm.RetrieveOneAs[offers.OfferOffer, offerProjection](
		ctx,
		repo,
		sq.Eq{"id": offerID},
	)
	if err != nil {
		t.Fatalf("retrieve offer projection: %v", err)
	}
	if projection.Status != "legal_review" || projection.StartAt.Status != pgtype.Null {
		t.Fatalf("unexpected enum or nullable timestamp: %#v", projection)
	}
	if len(projection.InvestorEligibility) != 2 || len(projection.MarketedJurisdictions) != 2 {
		t.Fatalf("unexpected JSON arrays: %#v", projection)
	}
	if projection.Data["nc_stamping_text"] != "Stamped \"quoted\" ✓" {
		t.Fatalf("unexpected JSON data: %#v", projection.Data)
	}
	if projection.RiskDisclosures != "Risk disclosure with Unicode ✓" {
		t.Fatalf("unexpected risk disclosure: %q", projection.RiskDisclosures)
	}
	if projection.TotalShares != "125.250000000000000000" ||
		projection.PricePerShare != "2.500000000000000000" {
		t.Fatalf(
			"fractional offer values did not retain exact PostgreSQL decimals: shares=%q price=%q",
			projection.TotalShares,
			projection.PricePerShare,
		)
	}
	if projection.FundStructure != "open_ended" {
		t.Fatalf("unexpected default fund structure: %q", projection.FundStructure)
	}
	if projection.DataRoomGroupID != nil {
		t.Fatalf("legacy offer unexpectedly has data-room binding: %v", *projection.DataRoomGroupID)
	}

	empty, err := orm.RetrieveAllAs[offers.OfferOffer, offerProjection](
		ctx,
		repo,
		sq.Eq{"id": -1},
	)
	if err != nil {
		t.Fatalf("retrieve empty offer projections: %v", err)
	}
	if empty == nil || len(empty) != 0 {
		t.Fatalf("expected non-nil empty projections, got %#v", empty)
	}

	var profileID int
	err = tx.QueryRow(ctx, `
		INSERT INTO investment_profiles(id, user_id)
		VALUES (-700001, $1)
		RETURNING id
	`, userID).Scan(&profileID)
	if err != nil {
		t.Fatalf("insert projection profile: %v", err)
	}

	var investmentID int
	err = tx.QueryRow(ctx, `
		INSERT INTO investment_investments(
			id, user_id, offer_id, profile_id, amount, price_per_share,
			number_of_shares
		) VALUES (
			-700001, $1, $2, $3, 123.456789, 2.5, 49.3827156
		)
		RETURNING id
	`, userID, offerID, profileID).Scan(&investmentID)
	if err != nil {
		t.Fatalf("insert projection investment: %v", err)
	}

	investment, err := orm.RetrieveOne[
		investments.InvestmentInvestment,
		*investments.InvestmentInvestment,
	](ctx, repo, sq.Eq{"id": investmentID})
	if err != nil {
		t.Fatalf("retrieve investment projection: %v", err)
	}
	if investment.Amount == nil ||
		*investment.Amount != "123.456789000000000000" ||
		investment.PricePerShare != "2.500000000000000000" ||
		investment.NumberOfShares != "49.382715600000000000" {
		t.Fatalf(
			"investment decimals were not projected exactly: amount=%v price=%q shares=%q",
			investment.Amount,
			investment.PricePerShare,
			investment.NumberOfShares,
		)
	}

	const externalStrategyAssetsRaw = "12345678901234567890123456789012345678901234567890"
	var navID int64
	err = tx.QueryRow(ctx, `
		INSERT INTO fund_nav_records(
			id, offer_id, nav, recorded_at, external_strategy_assets_raw
		) VALUES (
			-700001, $1, 1.234567890123456789, clock_timestamp(), $2
		)
		RETURNING id
	`, offerID, externalStrategyAssetsRaw).Scan(&navID)
	if err != nil {
		t.Fatalf("insert projection NAV: %v", err)
	}

	nav, err := orm.RetrieveOne[
		fundnavrecords.FundNAVRecord,
		*fundnavrecords.FundNAVRecord,
	](ctx, repo, sq.Eq{"id": navID})
	if err != nil {
		t.Fatalf("retrieve NAV projection: %v", err)
	}
	if nav.NAV != "1.234567890123456789" ||
		nav.ExternalStrategyAssetsRaw != externalStrategyAssetsRaw {
		t.Fatalf(
			"NAV values were not projected exactly: nav=%q strategy=%q",
			nav.NAV,
			nav.ExternalStrategyAssetsRaw,
		)
	}

	const claimableSharesRaw = "9876543210987654321098765432109876543210"
	var redemptionID int64
	err = tx.QueryRow(ctx, `
		INSERT INTO investment_redemptions(
			id, offer_id, profile_id, investment_id, vault_request_origin,
			idempotency_key, share_amount_raw, claimable_assets_raw,
			claimable_shares_raw, claimable_at
		) VALUES (
			-700001, $1, $2, $3, 'chain',
			'projection-redemption', $4, 0, $4, clock_timestamp()
		)
		RETURNING id
	`, offerID, profileID, investmentID, claimableSharesRaw).Scan(&redemptionID)
	if err != nil {
		t.Fatalf("insert projection redemption: %v", err)
	}

	redemption, err := orm.RetrieveOne[
		investments.InvestmentRedemption,
		*investments.InvestmentRedemption,
	](ctx, repo, sq.Eq{"id": redemptionID})
	if err != nil {
		t.Fatalf("retrieve redemption projection: %v", err)
	}
	if redemption.ShareAmountRaw != claimableSharesRaw ||
		redemption.ClaimableSharesRaw != claimableSharesRaw ||
		redemption.ClaimableAssetsRaw != "0" {
		t.Fatalf(
			"redemption exact/share-derived Claimable projection mismatch: %#v",
			redemption,
		)
	}
}
