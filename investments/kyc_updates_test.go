package investments

import (
	"reflect"
	"strings"
	"testing"

	"github.com/webdevelop-pro/i-models/profiles"
)

func TestMarkLegallyConfirmedQueryIsolatesIndividualProfile(t *testing.T) {
	t.Parallel()

	profile := &profiles.Profile{ID: 42, UserID: 7, Type: profiles.Individual}
	query, args := markLegallyConfirmedQuery(profile)

	for _, predicate := range []string{
		"ii.profile_id=ip.id",
		"ip.id=$2",
		"ii.user_id=$3",
		"ii.status=$4",
	} {
		if !strings.Contains(query, predicate) {
			t.Fatalf("query does not contain multi-profile isolation predicate %q: %s", predicate, query)
		}
	}

	wantArgs := []any{InvestmentTLegallyConfirmed, 42, 7, InvestmentTConfirmed}
	if !reflect.DeepEqual(args, wantArgs) {
		t.Fatalf("unexpected query arguments: got %#v, want %#v", args, wantArgs)
	}
}

func TestMarkLegallyConfirmedQueryUsesExactNonIndividualProfile(t *testing.T) {
	t.Parallel()

	profile := &profiles.Profile{ID: 51, UserID: 9, Type: profiles.Entity}
	query, args := markLegallyConfirmedQuery(profile)

	if !strings.Contains(query, "profile_id=$2") {
		t.Fatalf("query does not isolate the non-individual profile: %s", query)
	}
	wantArgs := []any{InvestmentTLegallyConfirmed, 51, InvestmentTConfirmed}
	if !reflect.DeepEqual(args, wantArgs) {
		t.Fatalf("unexpected query arguments: got %#v, want %#v", args, wantArgs)
	}
}
