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
	TotalShares           int                `db:"total_shares"`
	PricePerShare         float64            `db:"price_per_share"`
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

	var offerID int
	err = tx.QueryRow(ctx, `
		INSERT INTO offer_offers(
			name, slug, min_investment, status, data, investor_eligibility,
			marketed_jurisdictions, risk_disclosures, start_at, total_shares, price_per_share
		) VALUES (
			$1, $2, $3, $4, $5::jsonb, $6::jsonb, $7::jsonb, $8, NULL, $9, $10
		) RETURNING id
	`,
		"Projection offer",
		"projection-offer",
		100,
		"legal-review",
		`{"nc_stamping_text":"Stamped \"quoted\" ✓"}`,
		`["accredited","qualified_investor"]`,
		`["US","EU"]`,
		"Risk disclosure with Unicode ✓",
		125,
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
	if projection.Status != "legal-review" || projection.StartAt.Status != pgtype.Null {
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
	if got := float64(projection.TotalShares) * projection.PricePerShare; got != 312.5 {
		t.Fatalf("unexpected derived valuation: %v", got)
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
}
