// Package userinvitations models the shared user invitation aggregate.
package userinvitations

import "errors"

// errInvalidInvitationValue is the shared enum validation sentinel.
var errInvalidInvitationValue = errors.New("invitation value is not valid")

// Kind identifies the invitation workflow.
type Kind string

const (
	// KindTeam invites an operator to a tenant team.
	KindTeam Kind = "team"
	// KindInvestor invites an investor to create or join an investment profile.
	KindInvestor Kind = "investor"
)

// AllKinds returns every supported invitation kind.
func AllKinds() []Kind {
	return []Kind{KindTeam, KindInvestor}
}

// IsValid reports whether the kind is part of the invitation contract.
func (kind Kind) IsValid() error {
	switch kind {
	case KindTeam, KindInvestor:
		return nil
	default:
		return errInvalidInvitationValue
	}
}

// String returns the database representation of the kind.
func (kind Kind) String() string {
	return string(kind)
}

// Role identifies a team invitation's authorization role.
type Role string

const (
	// RoleOwner grants tenant-owner access.
	RoleOwner Role = "Owner"
	// RoleAdmin grants tenant-administrator access.
	RoleAdmin Role = "Admin"
	// RoleOps grants tenant-operations access.
	RoleOps Role = "Ops"
	// RoleAuditor grants read-oriented tenant-auditor access.
	RoleAuditor Role = "Auditor"
)

// AllRoles returns every supported team invitation role.
func AllRoles() []Role {
	return []Role{RoleOwner, RoleAdmin, RoleOps, RoleAuditor}
}

// IsValid reports whether the role is part of the invitation contract.
func (role Role) IsValid() error {
	switch role {
	case RoleOwner, RoleAdmin, RoleOps, RoleAuditor:
		return nil
	default:
		return errInvalidInvitationValue
	}
}

// String returns the database representation of the role.
func (role Role) String() string {
	return string(role)
}

// Status identifies the invitation lifecycle state.
type Status string

const (
	// StatusPending marks an invitation that can still be accepted.
	StatusPending Status = "pending"
	// StatusAccepted marks an invitation that has been consumed.
	StatusAccepted Status = "accepted"
	// StatusCancelled marks an invitation withdrawn by an operator.
	StatusCancelled Status = "cancelled"
	// StatusExpired marks an invitation whose acceptance window elapsed.
	StatusExpired Status = "expired"
)

// AllStatuses returns every supported invitation lifecycle state.
func AllStatuses() []Status {
	return []Status{StatusPending, StatusAccepted, StatusCancelled, StatusExpired}
}

// IsValid reports whether the status is part of the invitation contract.
func (status Status) IsValid() error {
	switch status {
	case StatusPending, StatusAccepted, StatusCancelled, StatusExpired:
		return nil
	default:
		return errInvalidInvitationValue
	}
}

// String returns the database representation of the status.
func (status Status) String() string {
	return string(status)
}

const (
	// AppLabel is the canonical Django permission app label.
	AppLabel = "user"
	// ModelName is the canonical Django permission model name.
	ModelName = "user_user_invitations"
	// TableName is the canonical PostgreSQL table name.
	TableName = "user_user_invitations"
)
