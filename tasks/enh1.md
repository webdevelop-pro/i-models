# i-models Query/Error Refactor and Model Consistency Fixes

## Summary
- Make `i-models` a neutral model/persistence component with db-agnostic not-found handling.
- Change generic read helpers to use `sq.Sqlizer`, then migrate `../evm-api` call sites in the same breaking change.
- Fix the current high-risk model-library issues: stale root `main.go`, inconsistent `Fields()`, and broken `ToJSON/ToMap` map generation.

## Flow

- Work in iterations.
- After each implementation iteration, run unit tests and required validation checks. Do not continue until tests are green, unless a failing test is explicitly documented as unrelated, flaky, or blocked with evidence.
- After every implementation iteration, create an independent review subagent using the golang-pro skills.
- The review subagent must inspect the changed code and the surrounding codebase independently.
- The review subagent must produce a written report containing:
  - Bugs or correctness issues
  - Missing requirements
  - Test gaps
  - Security or reliability concerns
  - Maintainability or architecture improvements
  - Any regressions or broken module boundaries
- The main agent must address every actionable critique from the review report.
- After fixes are made, run tests again and create a new independent review subagent.
- Repeat the review/fix/test loop until the review subagent reports no remaining actionable problems or improvements.
- The agent must not stop while there are unresolved actionable problems or improvements.
- If an issue cannot be addressed, the agent must document:
  - The issue
  - Why it cannot be fixed now
  - What evidence supports that decision
  - Whether it is blocked, out of scope, or requires user input

## Final Verification

- Once implementation appears complete, create a new independent verification subagent.
- The verification subagent must read `tasks/enh1.md`.
- The verification subagent must compare the implemented code against the task requirements.
- The verification subagent must produce a final verification report listing:
  - Fully implemented requirements
  - Partially implemented requirements
  - Missing requirements
  - Incorrect implementations
  - Additional risks or regressions
- The main agent must address every issue found in the final verification report.
- After addressing issues, run tests again and repeat final verification with a new independent subagent.
- Continue until the final verification subagent confirms that `tasks/enh1.md` has been correctly implemented and there are no remaining actionable issues.
- If the same issue remains unresolved after multiple attempts, the agent must not silently stop. It must escalate by documenting the blocker, the attempted fixes, and the exact remaining problem.


## Key Changes
- Query API:
  - Add `RetrieveOne[T](ctx, repo, where sq.Sqlizer, suffixes ...sq.Sqlizer)` and `RetrieveAll[T](ctx, repo, where sq.Sqlizer, suffixes ...sq.Sqlizer)`.
  - Superseded by `go-common/tasks/orm.md`: remove `RetriveOne/RetriveAll` wrappers instead of keeping deprecated aliases.
  - Apply `where` with Squirrel `Where`; apply trailing clauses with `SuffixExpr`, e.g. `sq.Expr("ORDER BY id LIMIT 1")`.
- Errors:
  - Add `models.ErrRecordNotFound`, wrapping `pgx.ErrNoRows` so both sentinels work during migration.
  - Check `pg.Query` errors before pgx collection to prevent nil-row panics.
  - Replace generic helper string-error construction with neutral wrapped errors.
- Safety:
  - Guard `Update` and `Delete` against empty predicates.
  - Keep existing `Update/Exists/Delete` `sq.Sqlizer` predicate support.

## Model Fixes
- Remove the stale root `main.go`; this is a library module and should not expose a broken root `package main`.
- Standardize `Fields()` across runtime model packages:
  - `Fields()` must return database select columns, not Go field names.
  - Prefer `db` tags plus `models.DefaultFields`; keep explicit lists only for real SQL expressions or aliases.
  - Fix affected packages such as filers, historylogs, pubsublogs, logs, offers relation models, and distributions relation/report models.
- Fix `ToJSON/ToMap` generation:
  - Methods that iterate DB-column names must call `GetValueByTag`, not `GetField`.
  - Apply this to wallet/EVM/transaction models and any same-pattern packages found in the audit.
  - Add serialization tests so non-empty sample structs produce expected DB-keyed maps.
- Exclude secrets from serialization:
  - EVM private keys and user passwords must not appear in JSON/YAML or `ToJSON`.


## Test Plan
- `i-models`:
  - Add tests for `RetrieveOne/RetrieveAll` SQL generation, suffix handling, query errors, and not-found wrapping.
  - Add tests for empty-predicate `Update/Delete` rejection.
  - Add `Fields()` tests for representative fixed packages.
  - Add `ToJSON/ToMap` tests for wallet, EVM wallet, and transaction models.
  - Run `go test ./...`.
- `evm-api`:
  - Update compile errors from new helper signatures.
  - Run `go test ./...`.
  - Verify representative not-found flows still return the same HTTP/domain behavior.

## Assumptions
- This is a coordinated breaking change across `i-models` and `evm-api`.
- The chosen db-agnostic sentinel name is `ErrRecordNotFound`.
- `pgx.ErrNoRows` compatibility is retained temporarily for migration safety.
