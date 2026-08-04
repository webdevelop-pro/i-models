package userinvitations

import "testing"

func TestInvitationEnums(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		values  []string
		invalid func(string) error
	}{
		{
			name:   "kinds",
			values: stringValues(AllKinds()),
			invalid: func(value string) error {
				return Kind(value).IsValid()
			},
		},
		{
			name:   "roles",
			values: stringValues(AllRoles()),
			invalid: func(value string) error {
				return Role(value).IsValid()
			},
		},
		{
			name:   "statuses",
			values: stringValues(AllStatuses()),
			invalid: func(value string) error {
				return Status(value).IsValid()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for _, value := range test.values {
				if err := test.invalid(value); err != nil {
					t.Fatalf("listed value %q is invalid: %v", value, err)
				}
			}
			if err := test.invalid("not-valid"); err == nil {
				t.Fatal("unknown value must be rejected")
			}
		})
	}
}

func TestUserInvitationProjectionOmitsCodeHash(t *testing.T) {
	t.Parallel()

	invitation := UserInvitation{ID: 7, CodeHash: "secret-digest"}
	payload := invitation.ToJSON()
	if payload["id"] != 7 {
		t.Fatalf("unexpected id: %#v", payload["id"])
	}
	if _, ok := payload["code_hash"]; ok {
		t.Fatal("code hash must not be exposed")
	}
}

func stringValues[T ~string](values []T) []string {
	result := make([]string, len(values))
	for index, value := range values {
		result[index] = string(value)
	}
	return result
}
