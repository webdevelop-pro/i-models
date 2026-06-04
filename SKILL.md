---
name: i-models
description: Shared Webdevelop Pro investment platform Go model library guidance. Use when working with github.com/webdevelop-pro/i-models packages for database model structs, Squirrel/pgx generic CRUD helpers, enum/null enum types, wallet/investment/offer/profile/filer/user models, Pub/Sub activity deduplication, PostgreSQL log models, local pgtype timestamp/json/null types, SQLBoiler template experiments, or tests involving i-models fixtures and repositories.
---

# I Models

Use this skill to work with `github.com/webdevelop-pro/i-models`, the shared Go
model library for investment-platform services. The root module is
`github.com/webdevelop-pro/i-models`; import package-level models directly, for
example `github.com/webdevelop-pro/i-models/transactions`.

Before assuming a model method exists, inspect that package's `models.go`.
Model packages are not perfectly uniform: some expose package-level `Get` or
`GetByID`, some only implement the shared model interface, and wallet/transaction
models have custom `Save` behavior.

Run validation from the repository root:

```bash
cd ../i-models
go test ./...
```

For template work only, use the separate `experiments` module:

```bash
cd ../i-models/experiments
go test ./...
```

## Package Map

- `models`: shared interfaces and generic CRUD helpers around Squirrel and pgx.
- `pgtype`: local trimmed copy of pgtype, kept to avoid vulnerable old deps while
  preserving database type mappings such as `Timestamptz`, `JSONB`, UUID, arrays,
  ranges, and `zeronull`.
- `transactions`, `wallets`, `fundingsources`: wallet and transaction models,
  enums, and custom post-update side effects.
- `investments`, `profiles`, `offers`, `distributions`: investment domain
  models, enum/null enum types, comments, filers, reports, KYC/accreditation.
- `users`, `emails`, `notifications`, `filers`, `historylogs`: application
  support tables and model structs.
- `evmwallets`, `evmtransfers`, `evmcontracts`, `evmchainaccounts`,
  `evmwalletbalances`, `evmwalletoperations`: EVM domain models and enums.
- `logs`: HTTP request/response log model and helpers for `django_content_type`.
- `pubsublogs`: stored Pub/Sub message log model.
- `pubsubactivities`: message-level Pub/Sub dedup store compatible with
  `go-common/queue.Deduper`.
- `experiments`: SQLBoiler template experiments; do not treat it as part of the
  runtime module.

## Core Interfaces

Generic helpers expect model structs to expose table and field metadata.

```go
type Repository interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Lg() logger.Logger
}

type Model interface {
	SetID(any)
	GetID() any
	ToJSON() map[string]any
	Fields() []string
	Table() string
}
```

Use a `go-common/db.DB` or anything implementing `db.Repository`/`models.Repository`
as the repository argument.

## Generic Queries

The shared query helper names intentionally use the current misspelling:
`RetriveOne` and `RetriveAll`. Use those exact names unless renaming the API
across all consumers.

```go
file, err := models.RetriveOne[filers.FilerFiler](
	ctx,
	repo,
	map[string]any{"id": fileID},
)
if err != nil {
	return nil, err
}
```

Use `RetriveAll` for lists. Extra string parameters are appended to the SQL, so
only pass trusted static fragments such as ordering or limits.

```go
files, err := models.RetriveAll[filers.FilerFiler](
	ctx,
	repo,
	map[string]any{"user_id": userID},
	"ORDER BY created_at DESC",
	"LIMIT 50",
)
```

Use `Create` for inserts. It appends `RETURNING id`, scans the generated id,
sets it on the model, and injects the repository through `SetDB`.

```go
notification, err := models.Create[notifications.Notification](
	ctx,
	repo,
	map[string]any{
		"user_id": userID,
		"content": "Wallet updated",
		"data": map[string]any{
			"wallet_id": walletID,
		},
	},
)
```

`Update`, `Exists`, and `Delete` accept a `where map[string]any` plus optional
Squirrel expressions for non-map predicates.

```go
updated, err := models.Update[transactions.Transaction](
	ctx,
	repo,
	map[string]any{"id": txID},
	map[string]any{"status": transactions.TransactionsStatusTProcessed},
	sq.Expr("updated_at < now()"),
)
if err != nil {
	return err
}
if !updated {
	return models.ErrNotFound
}
```

```go
exists, err := models.Exists[offers.OfferOffer](
	ctx,
	repo,
	map[string]any{"slug": slug},
	sq.NotEq{"id": currentOfferID},
)
```

For package-specific helpers, prefer them when present because they restore
model-internal state such as `db` and `fns`.

```go
wallet, err := wallets.GetByID(ctx, repo, walletID)
tx, err := transactions.Get(ctx, repo, map[string]any{"entity_id": entityID})
offer, err := offers.Get(ctx, repo, map[string]any{"slug": slug})
```

Use `errors.Is(err, pgx.ErrNoRows)` for not-found checks from retrieval helpers;
`ErrNotFound` wraps `pgx.ErrNoRows` for stack preservation.

## Model Mutations

Some models have setter methods that append to `updatedFields`. Use setters
before `Save`; direct field mutation will not be persisted by custom `Save`
methods unless the field is also in `updatedFields`.

```go
tx, err := transactions.GetByID(ctx, repo, txID)
if err != nil {
	return err
}

tx.SetStatus(transactions.TransactionsStatusTProcessed)
if err := tx.Save(ctx, publishEvent); err != nil {
	return err
}
```

`Save` methods generally:

- require a non-zero `ID`;
- call `models.Update` with only tracked fields;
- emit a `pclient.PostUpdate` event through the provided callback;
- may create notifications or history log rows as side effects.

For plain models without custom setters or `Save`, use `models.Update` directly.

## Enums And Null Enums

Enum packages use string aliases with constants, `All...()` helpers, `IsValid`,
and `String`.

```go
status := transactions.TransactionsStatusTProcessed
if err := status.IsValid(); err != nil {
	return err
}
```

Several SQL-nullable enums use `Null...` wrappers with `Valid`, `SetValid`,
`Scan`, `Value`, `MarshalJSON`, and `UnmarshalJSON`, especially under `offers`
and `distributions`.

```go
offer.Status = offers.NewNullOfferT(offers.OfferTPublished, true)
if offer.Status.Valid {
	status := offer.Status.Val.String()
	_ = status
}
```

Use existing enum constants rather than raw strings in service code. When adding
new enum values, update the constants, `All...()` list, `IsValid`, and any null
enum scanner/marshaler logic in the same package.

## PostgreSQL Types

Use `github.com/webdevelop-pro/i-models/pgtype` types when a model field already
uses them. Most timestamp columns use `pgtype.Timestamptz`; JSON/JSONB and array
columns may use local pgtype wrappers. Do not replace these with `time.Time`,
`map[string]any`, or pgx native types unless every scanner/encoder consumer is
updated.

```go
type Row struct {
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}
```

`pgtype` is a trimmed local copy kept to avoid vulnerable dependencies in the
original package. Edit it only for database type mapping work.

## Pub/Sub Dedup

Use `pubsubactivities` as the database-backed implementation for
`go-common/queue.Deduper`.

```go
type Deduper struct {
	repo db.Repository
}

func (d Deduper) Claim(ctx context.Context, service, topic, msgID string, attempt int) (bool, error) {
	return pubsubactivities.Claim(ctx, d.repo, msgID, service, topic, attempt)
}

func (d Deduper) MarkProcessed(ctx context.Context, service, msgID string) error {
	return pubsubactivities.MarkProcessed(ctx, d.repo, msgID, service)
}

func (d Deduper) MarkFailed(ctx context.Context, service, msgID, lastErr string) error {
	return pubsubactivities.MarkFailed(ctx, d.repo, msgID, service, lastErr)
}
```

In pull-mode worker tests using the Pub/Sub emulator, clear
`pubsub_activities`; emulator message ids can restart at `1`.

```go
fixtures := []tests.FixturesManager{
	dbtests.NewFixturesManager(ctx, dbtests.NewFixture("files", "fixtures/files.json")),
	qtests.NewFixturesManager(ctx, qtests.NewFixture(topic, subscription, "")),
	pubsubactivities.NewCleanFixtures(ctx),
}
```

Or add the clear action as the first scenario action:

```go
TestActions: []tests.SomeAction{
	pubsubactivities.ClearAction(),
	// publish event...
}
```

## HTTP And Service Logs

Use `logs.LogHttpRequest` when a service needs to persist third-party or
incoming/outgoing HTTP logs in the `logs` table. It reads `keys.RequestID` and
`keys.MSGID` from context and looks up the Django content type by app/model.

```go
logEntry, err := logs.LogHttpRequest(
	ctx,
	repo,
	logs.LogTypeTOutcoming,
	objectID,
	logs.ServicesTDwolla,
	transactions.AppLabel,
	transactions.ModelName,
	req,
	rawBody,
)
```

For client hooks, use `LogLog.LogRequest` and `LogLog.LogResponse`; they buffer
and restore request/response bodies after reading.

## Working On Models

When adding or updating a model package:

- keep field tags aligned with database columns and JSON output;
- implement `Fields`, `Table`, `GetID`, and `SetID`;
- implement `SetDB` if the model can be created through `models.Create` or has
  methods that need the repository later;
- use `models.DefaultFields(&model)` only when `db` tags are reliable;
- keep `ToJSON`/`ToMap` keys in API-facing snake_case unless the package already
  uses a different convention;
- add enum validation helpers when adding a string enum type;
- preserve existing `AppLabel`, `ModelName`, and `TableName` constants because
  history logs and content-type lookups depend on them.

For SQLBoiler template changes, work in `experiments/templates`, compare against
`experiments/templates_orig`, and run the separate `experiments` module tests if
you touch Go code there. The README command is:

```bash
sqlboiler crdb -c sqlboiler.toml --no-auto-timestamps --add-global-variants --no-tests --templates <path-to-templates>/templates/main -p models -o ./pkg/models
```
