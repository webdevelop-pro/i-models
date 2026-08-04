package users

import (
	"context"
	"errors"

	"github.com/global-torque/go-common/db/v2"
	"github.com/jackc/pgx/v5"

	"github.com/global-torque/go-common/orm/v2/pgtype"
)

const Table = "user_users"

// ToDo
// Add all fields
type UserData struct {
	NCIssuerID string `json:"nc_issuer_id"`
}

// UserUser is an object representing the database table.
type UserUser struct {
	ID          int                `db:"id" json:"id" yaml:"id"`
	Password    string             `db:"password" json:"-" yaml:"-"`
	Email       string             `db:"email" json:"email" yaml:"email"`
	FirstName   string             `db:"first_name" json:"first_name" yaml:"first_name"`
	LastName    string             `db:"last_name" json:"last_name" yaml:"last_name"`
	IsStaff     bool               `db:"is_staff" json:"is_staff" yaml:"is_staff"`
	IsSuperuser bool               `db:"is_superuser" json:"is_superuser" yaml:"is_superuser"`
	IsActive    bool               `db:"is_active" json:"is_active" yaml:"is_active"`
	IdentityID  string             `db:"identity_id" json:"identity_id" yaml:"identity_id"`
	SiteID      *int               `db:"site_id" json:"site_id,omitempty" yaml:"site_id,omitempty"`
	FacebookID  string             `db:"facebook_id" json:"facebook_id" yaml:"facebook_id"`
	LinkedinID  string             `db:"linkedin_id" json:"linkedin_id" yaml:"linkedin_id"`
	GoogleID    string             `db:"google_id" json:"google_id" yaml:"google_id"`
	Phone       string             `db:"phone" json:"phone" yaml:"phone"`
	IPAddress   string             `db:"ip_address" json:"ip_address" yaml:"ip_address"`
	UserAgent   string             `db:"user_agent" json:"user_agent" yaml:"user_agent"`
	Timezone    string             `db:"timezone" json:"timezone" yaml:"timezone"`
	Social      string             `db:"social" json:"social" yaml:"social"`
	Data        UserData           `db:"data" json:"data" yaml:"data"`
	LastLogin   pgtype.Timestamptz `db:"last_login" json:"last_login" yaml:"last_login"`
	CreatedAt   pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	ImageLinkID *int               `db:"image_link_id" json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`

	db db.Repository `db:"-" json:"-"`
}

func New(db db.Repository) *UserUser {
	return &UserUser{
		db: db,
	}
}

func (model UserUser) ToJSON() map[string]any {
	return map[string]any{
		"id":            model.ID,
		"email":         model.Email,
		"first_name":    model.FirstName,
		"last_name":     model.LastName,
		"is_staff":      model.IsStaff,
		"is_superuser":  model.IsSuperuser,
		"is_active":     model.IsActive,
		"identity_id":   model.IdentityID,
		"site_id":       model.SiteID,
		"facebook_id":   model.FacebookID,
		"linkedin_id":   model.LinkedinID,
		"google_id":     model.GoogleID,
		"phone":         model.Phone,
		"ip_address":    model.IPAddress,
		"user_agent":    model.UserAgent,
		"timezone":      model.Timezone,
		"social":        model.Social,
		"data":          model.Data,
		"last_login":    model.LastLogin,
		"created_at":    model.CreatedAt,
		"updated_at":    model.UpdatedAt,
		"image_link_id": model.ImageLinkID,
	}
}

func (user UserUser) Fields() []string {
	return []string{"id", "email", "first_name", "last_name", "identity_id", "site_id", "phone"}
}

func (model UserUser) Table() string {
	return Table
}

func (model UserUser) GetID() any {
	return model.ID
}

func (model *UserUser) SetID(id any) {
	model.ID = id.(int)
}

func (model *UserUser) SetDB(db db.Repository) {
	model.db = db
}

// HasGroup reports whether the user belongs to an auth_group by name.
func (model UserUser) HasGroup(ctx context.Context, repo db.Repository, groupName string) (bool, error) {
	if model.ID == 0 {
		return false, nil
	}

	var exists int
	err := repo.QueryRow(
		ctx,
		`
			SELECT 1
			FROM auth_group t1
			JOIN user_users_groups t2 ON t2.group_id = t1.id
			WHERE t2.account_id = $1 AND t1.name = $2
			LIMIT 1
		`,
		model.ID,
		groupName,
	).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
