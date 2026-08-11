# `investments`

This package is the shared investment-domain schema. Its main projections are:

- `InvestmentInvestment` (`investment_investments`) for subscriptions/funding;
- `InvestmentRedemption` (`investment_redemptions`) for redemption requests;
- `InvestmentProfile` (`investment_profiles`) for the investor-specific view
  used by investment and escrow flows.

It owns the canonical payment, escrow, funding, investment, redemption, vault,
profile, KYC, and accreditation enum families. `InvestmentProfile` is not a
replacement for [`profiles.Profile`](../profiles/README.md): both map the same
broad table boundary for different service projections.

## Concurrency helpers

`RetrieveInvestmentForUpdate` and `RetrieveRedemptionForUpdate` issue `FOR
UPDATE`; pass a transaction-scoped repository and keep the subsequent write in
that transaction. `UpdateUnlockedInvestment` and `UpdateUnlockedRedemption`
perform compare-and-swap updates against the expected state and request lock.
Always check their returned boolean.

```go
investment, err := investments.RetrieveInvestmentForUpdate(ctx, tx, id)
// validate the locked row
updated, err := investments.UpdateUnlockedInvestment(
    ctx, tx, id, investment.Status,
    map[string]any{"status": investments.InvestmentTClosedSuccessfully},
)
```

`MarkLegallyConfirmedForProfile` advances confirmed investments after profile
KYC/accreditation conditions are met.

## Consumers

`evm-api` and `payment-api` use the enums and locked transitions for crypto and
vault settlement. `wallet-api` uses investment/funding states for direct
funding. `escrow-api`, `kyc-api`, and `esign-api` use investment/profile
projections. `email-worker` uses them to qualify and render domain-event mail.

See the [root guide](../README.md) for transaction rules.
