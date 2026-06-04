## evm-api Migration
- Update all `../evm-api` generic helper calls:
  - `map[string]any{...}` to `sq.Eq{...}`.
  - string suffixes to `sq.Expr("ORDER BY ...")`.
  - helper-origin not-found checks to `errors.Is(err, models.ErrRecordNotFound)`.
- Keep `pgx.ErrNoRows` checks only for raw `QueryRow().Scan` paths.
- Keep evm-api’s narrow validator structs; they are safer than exposing broad generated model structs.