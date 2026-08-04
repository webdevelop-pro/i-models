package userinvitations

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/webdevelop-pro/i-models/investments"
)

// UserInvitation is the persisted team or investor invitation aggregate.
//
//nolint:recvcheck // Persistence interfaces require value readers and pointer mutators.
type UserInvitation struct {
	ID                int                   `db:"id" json:"id" yaml:"id"`
	SiteID            int                   `db:"site_id" json:"site_id" yaml:"site_id"`
	Email             string                `db:"email" json:"email" yaml:"email"`
	FirstName         string                `db:"first_name" json:"first_name" yaml:"first_name"`
	LastName          string                `db:"last_name" json:"last_name" yaml:"last_name"`
	Kind              Kind                  `db:"kind" json:"kind" yaml:"kind"`
	Role              *Role                 `db:"role" json:"role,omitempty" yaml:"role,omitempty"`
	ProfileType       *investments.ProfileT `db:"profile_type" json:"profile_type,omitempty" yaml:"profile_type,omitempty"`
	FormTemplateID    *string               `db:"form_template_id" json:"form_template_id,omitempty" yaml:"form_template_id,omitempty"` //nolint:lll // Tags mirror the column name.
	FundID            *int                  `db:"fund_id" json:"fund_id,omitempty" yaml:"fund_id,omitempty"`
	FundName          *string               `db:"fund_name" json:"fund_name,omitempty" yaml:"fund_name,omitempty"`
	CodeHash          string                `db:"code_hash" json:"-" yaml:"-"`
	Status            Status                `db:"status" json:"status" yaml:"status"`
	ExpiresAt         pgtype.Timestamptz    `db:"expires_at" json:"expires_at" yaml:"expires_at"`
	InvitedByID       *int                  `db:"invited_by_id" json:"invited_by_id,omitempty" yaml:"invited_by_id,omitempty"`                   //nolint:lll // Tags mirror the column name.
	AcceptedUserID    *int                  `db:"accepted_user_id" json:"accepted_user_id,omitempty" yaml:"accepted_user_id,omitempty"`          //nolint:lll // Tags mirror the column name.
	AcceptedProfileID *int                  `db:"accepted_profile_id" json:"accepted_profile_id,omitempty" yaml:"accepted_profile_id,omitempty"` //nolint:lll // Tags mirror the column name.
	AcceptedAt        pgtype.Timestamptz    `db:"accepted_at" json:"accepted_at" yaml:"accepted_at"`
	CancelledAt       pgtype.Timestamptz    `db:"cancelled_at" json:"cancelled_at" yaml:"cancelled_at"`
	LastSentAt        pgtype.Timestamptz    `db:"last_sent_at" json:"last_sent_at" yaml:"last_sent_at"`
	CreatedAt         pgtype.Timestamptz    `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt         pgtype.Timestamptz    `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	db db.Repository `db:"-" json:"-" yaml:"-"`
}

// New creates a user invitation bound to a repository.
func New(repo db.Repository) *UserInvitation {
	return &UserInvitation{db: repo}
}

// ToJSON returns the public database-field projection without the code hash.
func (invitation UserInvitation) ToJSON() map[string]any {
	return map[string]any{
		"id":                  invitation.ID,
		"site_id":             invitation.SiteID,
		"email":               invitation.Email,
		"first_name":          invitation.FirstName,
		"last_name":           invitation.LastName,
		"kind":                invitation.Kind,
		"role":                invitation.Role,
		"profile_type":        invitation.ProfileType,
		"form_template_id":    invitation.FormTemplateID,
		"fund_id":             invitation.FundID,
		"fund_name":           invitation.FundName,
		"status":              invitation.Status,
		"expires_at":          invitation.ExpiresAt,
		"invited_by_id":       invitation.InvitedByID,
		"accepted_user_id":    invitation.AcceptedUserID,
		"accepted_profile_id": invitation.AcceptedProfileID,
		"accepted_at":         invitation.AcceptedAt,
		"cancelled_at":        invitation.CancelledAt,
		"last_sent_at":        invitation.LastSentAt,
		"created_at":          invitation.CreatedAt,
		"updated_at":          invitation.UpdatedAt,
	}
}

// Fields returns every persisted column required to hydrate an invitation.
func (invitation UserInvitation) Fields() []string {
	return []string{
		"id",
		"site_id",
		"email",
		"first_name",
		"last_name",
		"kind",
		"role",
		"profile_type",
		"form_template_id",
		"fund_id",
		"fund_name",
		"code_hash",
		"status",
		"expires_at",
		"invited_by_id",
		"accepted_user_id",
		"accepted_profile_id",
		"accepted_at",
		"cancelled_at",
		"last_sent_at",
		"created_at",
		"updated_at",
	}
}

// Table returns the canonical invitation table name.
func (invitation UserInvitation) Table() string {
	return TableName
}

// GetID returns the invitation primary key.
func (invitation UserInvitation) GetID() any {
	return invitation.ID
}

// SetID assigns the invitation primary key.
func (invitation *UserInvitation) SetID(id any) {
	value, ok := id.(int)
	if !ok {
		panic("user invitation ID must be an int")
	}

	invitation.ID = value
}

// SetDB binds the invitation to a repository.
func (invitation *UserInvitation) SetDB(repo db.Repository) {
	invitation.db = repo
}
