package pubsubactivities

import (
	"context"
	"errors"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	testEventID   = "b0000000-0000-4000-8000-000000000001"
	testOtherID   = "b0000000-0000-4000-8000-000000000002"
	testClaim     = "6d0aaf23-d5ea-4ed5-b020-60fb9ba72155"
	testService   = "wallet-api-domain-events"
	testTopic     = "domain-events"
	testMessageID = "9812345678901234"
)

type reactionRepo struct {
	tx       *reactionTx
	beginErr error
	begun    int
}

func (repo *reactionRepo) Begin(context.Context) (pgx.Tx, error) {
	repo.begun++
	if repo.beginErr != nil {
		return nil, repo.beginErr
	}
	return repo.tx, nil
}

func (*reactionRepo) Exec(context.Context, string, ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, errors.New("unexpected repository Exec")
}

func (*reactionRepo) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected repository Query")
}

func (*reactionRepo) QueryRow(context.Context, string, ...any) pgx.Row {
	return rowFunc(func(...any) error { return errors.New("unexpected repository QueryRow") })
}

type reactionTx struct {
	rows       []pgx.Row
	queries    []string
	args       [][]any
	committed  bool
	rolledBack bool
	commitErr  error
}

func (tx *reactionTx) Begin(context.Context) (pgx.Tx, error) { return tx, nil }

func (tx *reactionTx) Commit(context.Context) error {
	tx.committed = true
	return tx.commitErr
}

func (tx *reactionTx) Rollback(context.Context) error {
	tx.rolledBack = true
	return nil
}

func (*reactionTx) CopyFrom(
	context.Context,
	pgx.Identifier,
	[]string,
	pgx.CopyFromSource,
) (int64, error) {
	return 0, errors.New("unexpected CopyFrom")
}

func (*reactionTx) SendBatch(context.Context, *pgx.Batch) pgx.BatchResults { return nil }
func (*reactionTx) LargeObjects() pgx.LargeObjects                         { return pgx.LargeObjects{} }
func (*reactionTx) Prepare(
	context.Context,
	string,
	string,
) (*pgconn.StatementDescription, error) {
	return nil, errors.New("unexpected Prepare")
}

func (*reactionTx) Exec(context.Context, string, ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, errors.New("unexpected Exec")
}

func (*reactionTx) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (tx *reactionTx) QueryRow(_ context.Context, query string, args ...any) pgx.Row {
	tx.queries = append(tx.queries, query)
	tx.args = append(tx.args, args)
	if len(tx.rows) == 0 {
		return rowFunc(func(...any) error { return errors.New("unexpected QueryRow") })
	}
	row := tx.rows[0]
	tx.rows = tx.rows[1:]
	return row
}

func (*reactionTx) Conn() *pgx.Conn { return nil }

func scanString(value string) pgx.Row {
	return rowFunc(func(dest ...any) error {
		ptr, ok := dest[0].(*string)
		if !ok {
			return errors.New("destination is not *string")
		}
		*ptr = value
		return nil
	})
}

func scanStatusAndDelivery(status string, hasDelivery bool) pgx.Row {
	return rowFunc(func(dest ...any) error {
		statusPtr, ok := dest[0].(*string)
		if !ok {
			return errors.New("status destination is not *string")
		}
		deliveryPtr, ok := dest[1].(*bool)
		if !ok {
			return errors.New("delivery destination is not *bool")
		}
		*statusPtr = status
		*deliveryPtr = hasDelivery
		return nil
	})
}

func TestStatusContract(t *testing.T) {
	t.Parallel()

	terminal := []Status{StatusProcessed, StatusReacted, StatusRejected, StatusResolved}
	for _, status := range terminal {
		if !status.IsValid() || !status.IsTerminal() {
			t.Errorf("status %q should be valid and terminal", status)
		}
	}
	for _, status := range []Status{StatusProcessing, StatusFailed} {
		if !status.IsValid() || status.IsTerminal() {
			t.Errorf("status %q should be valid and non-terminal", status)
		}
	}
	if Status("future").IsValid() || Status("future").IsTerminal() {
		t.Fatal("unknown status must be invalid and non-terminal")
	}
}

func TestClaimLeaseTreatsEveryTerminalStatusAsDuplicate(t *testing.T) {
	t.Parallel()

	for _, status := range []Status{StatusProcessed, StatusReacted, StatusRejected, StatusResolved} {
		status := status
		t.Run(string(status), func(t *testing.T) {
			t.Parallel()
			repo := &claimRepository{queryRows: []pgx.Row{
				rowFunc(func(...any) error { return pgx.ErrNoRows }),
				scanString(string(status)),
			}}
			token, claimed, err := ClaimLease(
				context.Background(), repo, testEventID, testService, testTopic, 1,
			)
			if err != nil || claimed || token != "" {
				t.Fatalf("ClaimLease() token=%q claimed=%v err=%v", token, claimed, err)
			}
		})
	}
}

func TestClaimLeaseRejectsUnknownStatus(t *testing.T) {
	t.Parallel()

	repo := &claimRepository{queryRows: []pgx.Row{
		rowFunc(func(...any) error { return pgx.ErrNoRows }),
		scanString("future"),
	}}
	_, _, err := ClaimLease(context.Background(), repo, testEventID, testService, testTopic, 1)
	if !errors.Is(err, ErrUnknownStatus) {
		t.Fatalf("ClaimLease() error = %v, want ErrUnknownStatus", err)
	}
}

func TestClaimDomainReactionCommitsClaimAndDelivery(t *testing.T) {
	t.Parallel()

	tx := &reactionTx{rows: []pgx.Row{scanString(testClaim), scanString(testEventID)}}
	repo := &reactionRepo{tx: tx}
	token, claimed, err := ClaimDomainReaction(
		context.Background(), repo, testEventID, testService, testTopic, testMessageID, 3,
	)
	if err != nil || !claimed || token != testClaim {
		t.Fatalf("ClaimDomainReaction() token=%q claimed=%v err=%v", token, claimed, err)
	}
	if !tx.committed {
		t.Fatal("transaction was not committed")
	}
	if len(tx.queries) != 2 || !strings.Contains(tx.queries[1], "pubsub_domain_event_deliveries") {
		t.Fatalf("queries = %#v, want claim followed by delivery", tx.queries)
	}
	if !strings.Contains(tx.queries[0], "last_error  = ''") {
		t.Fatal("claim query does not clear stale last_error")
	}
}

func TestClaimDomainReactionTerminalDuplicateDoesNotMutateDelivery(t *testing.T) {
	t.Parallel()

	tx := &reactionTx{rows: []pgx.Row{
		rowFunc(func(...any) error { return pgx.ErrNoRows }),
		scanString(string(StatusReacted)),
	}}
	token, claimed, err := ClaimDomainReaction(
		context.Background(), &reactionRepo{tx: tx},
		testEventID, testService, testTopic, "republished-message", 9,
	)
	if err != nil || claimed || token != "" {
		t.Fatalf("ClaimDomainReaction() token=%q claimed=%v err=%v", token, claimed, err)
	}
	if tx.committed || len(tx.queries) != 2 {
		t.Fatalf("terminal duplicate committed=%v queries=%d", tx.committed, len(tx.queries))
	}
}

func TestClaimDomainReactionRollsBackDeliveryConflict(t *testing.T) {
	t.Parallel()

	tx := &reactionTx{rows: []pgx.Row{
		scanString(testClaim),
		rowFunc(func(...any) error { return pgx.ErrNoRows }),
	}}
	_, claimed, err := ClaimDomainReaction(
		context.Background(), &reactionRepo{tx: tx},
		testEventID, testService, testTopic, testMessageID, 1,
	)
	if !errors.Is(err, ErrDeliveryIdentityConflict) || claimed {
		t.Fatalf("ClaimDomainReaction() claimed=%v err=%v", claimed, err)
	}
	if tx.committed || !tx.rolledBack {
		t.Fatalf("conflict committed=%v rolledBack=%v", tx.committed, tx.rolledBack)
	}
}

func TestClaimDomainReactionValidatesIdentityBeforeBegin(t *testing.T) {
	t.Parallel()

	repo := &reactionRepo{}
	_, _, err := ClaimDomainReaction(
		context.Background(), repo, "not-a-uuid", testService, testTopic, testMessageID, 1,
	)
	if !errors.Is(err, ErrInvalidIdentity) || repo.begun != 0 {
		t.Fatalf("ClaimDomainReaction() begun=%d err=%v", repo.begun, err)
	}
}

func TestResolveDomainReactionStates(t *testing.T) {
	t.Parallel()

	t.Run("resolves prior failed delivery", func(t *testing.T) {
		t.Parallel()
		tx := &reactionTx{rows: []pgx.Row{scanString(string(StatusResolved))}}
		resolved, err := ResolveDomainReaction(context.Background(), &reactionRepo{tx: tx}, testEventID, testService)
		if err != nil || !resolved || !tx.committed {
			t.Fatalf("ResolveDomainReaction() resolved=%v committed=%v err=%v", resolved, tx.committed, err)
		}
	})

	t.Run("first ignore writes nothing", func(t *testing.T) {
		t.Parallel()
		tx := &reactionTx{rows: []pgx.Row{
			rowFunc(func(...any) error { return pgx.ErrNoRows }),
			rowFunc(func(...any) error { return pgx.ErrNoRows }),
		}}
		resolved, err := ResolveDomainReaction(context.Background(), &reactionRepo{tx: tx}, testEventID, testService)
		if err != nil || resolved || tx.committed {
			t.Fatalf("ResolveDomainReaction() resolved=%v committed=%v err=%v", resolved, tx.committed, err)
		}
	})

	t.Run("live lease remains retryable", func(t *testing.T) {
		t.Parallel()
		tx := &reactionTx{rows: []pgx.Row{
			rowFunc(func(...any) error { return pgx.ErrNoRows }),
			scanStatusAndDelivery(string(StatusProcessing), true),
		}}
		resolved, err := ResolveDomainReaction(context.Background(), &reactionRepo{tx: tx}, testEventID, testService)
		if !errors.Is(err, ErrInProgress) || resolved || tx.committed {
			t.Fatalf("ResolveDomainReaction() resolved=%v committed=%v err=%v", resolved, tx.committed, err)
		}
	})
}

type rejectionRepo struct {
	query string
	args  []any
	row   pgx.Row
}

func (*rejectionRepo) Exec(context.Context, string, ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, errors.New("unexpected Exec")
}

func (*rejectionRepo) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (repo *rejectionRepo) QueryRow(_ context.Context, query string, args ...any) pgx.Row {
	repo.query, repo.args = query, args
	return repo.row
}

func TestRejectionsUseCorrectIdentityAndMonotonicUpsert(t *testing.T) {
	t.Parallel()

	transport := &rejectionRepo{row: scanString(string(StatusRejected))}
	err := RecordTransportRejection(
		context.Background(), transport, testMessageID, testService, testTopic, 4,
		strings.Repeat("x", maxLastError+10),
	)
	if err != nil {
		t.Fatalf("RecordTransportRejection() error = %v", err)
	}
	if transport.args[0] != transportPrefix+testMessageID || transport.args[3] != 4 {
		t.Fatalf("transport rejection args = %#v", transport.args)
	}
	if len(transport.args[4].(string)) != maxLastError {
		t.Fatalf("transport rejection error length = %d", len(transport.args[4].(string)))
	}
	if !strings.Contains(transport.query, "GREATEST(pubsub_activities.attempt, EXCLUDED.attempt)") {
		t.Fatal("rejection upsert is not monotonic")
	}

	application := &rejectionRepo{row: scanString(string(StatusRejected))}
	if err := RecordApplicationRejection(
		context.Background(), application, testEventID, testService, testTopic, 2, "invalid contract",
	); err != nil {
		t.Fatalf("RecordApplicationRejection() error = %v", err)
	}
	if application.args[0] != testEventID {
		t.Fatalf("application rejection identity = %v", application.args[0])
	}
}

type terminalContextRepo struct {
	ctxErr error
	tag    pgconn.CommandTag
}

func (repo *terminalContextRepo) Exec(ctx context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
	repo.ctxErr = ctx.Err()
	return repo.tag, nil
}

func (*terminalContextRepo) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (*terminalContextRepo) QueryRow(context.Context, string, ...any) pgx.Row {
	return rowFunc(func(...any) error { return errors.New("unexpected QueryRow") })
}

func TestDomainTerminalWritesDetachCancellationAndFence(t *testing.T) {
	t.Parallel()

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	repo := &terminalContextRepo{tag: pgconn.NewCommandTag("UPDATE 1")}
	if err := MarkReactedLease(cancelled, repo, testEventID, testService, testClaim); err != nil {
		t.Fatalf("MarkReactedLease() error = %v", err)
	}
	if repo.ctxErr != nil {
		t.Fatalf("terminal context error = %v, want detached context", repo.ctxErr)
	}

	repo.tag = pgconn.NewCommandTag("UPDATE 0")
	err := MarkRejectedLease(cancelled, repo, testEventID, testService, testClaim, "permanent")
	if !errors.Is(err, ErrClaimLost) {
		t.Fatalf("MarkRejectedLease() error = %v, want ErrClaimLost", err)
	}
}

func TestTransportRejectionValidatesNamespacedLength(t *testing.T) {
	t.Parallel()

	repo := &rejectionRepo{row: scanString(string(StatusRejected))}
	err := RecordTransportRejection(
		context.Background(), repo,
		strings.Repeat("m", maxMessageIDLength-len(transportPrefix)+1),
		testService, testTopic, 1, "invalid",
	)
	if !errors.Is(err, ErrInvalidIdentity) || repo.query != "" {
		t.Fatalf("RecordTransportRejection() query=%q err=%v", repo.query, err)
	}
}

func TestTruncateErrorPreservesValidUTF8(t *testing.T) {
	t.Parallel()

	got := truncateError(strings.Repeat("x", maxLastError-1) + "🙂")
	if len(got) > maxLastError || !utf8.ValidString(got) {
		t.Fatalf("truncateError() length=%d valid=%v", len(got), utf8.ValidString(got))
	}
}

func TestClaimDomainReactionDetectsUnexpectedReturnedIdentity(t *testing.T) {
	t.Parallel()

	tx := &reactionTx{rows: []pgx.Row{scanString(testClaim), scanString(testOtherID)}}
	_, claimed, err := ClaimDomainReaction(
		context.Background(), &reactionRepo{tx: tx},
		testEventID, testService, testTopic, testMessageID, 1,
	)
	if !errors.Is(err, ErrDeliveryIdentityConflict) || claimed || tx.committed {
		t.Fatalf("ClaimDomainReaction() claimed=%v committed=%v err=%v", claimed, tx.committed, err)
	}
}
