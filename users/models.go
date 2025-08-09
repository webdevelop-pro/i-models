package users

import (
	"github.com/webdevelop-pro/i-models/models"

	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/pgtype"
)

const Table = "user_users"

// ToDo
// Add all fields
type UserData struct {
	NCIssuerID string `json:"nc_issuer_id"`
}

// UserUser is an object representing the database table.
type UserUser struct {
	ID          int                `db:"-" json:"id" yaml:"id"`
	Password    string             `db:"-" json:"password" yaml:"password"`
	Email       string             `db:"-" json:"email" yaml:"email"`
	FirstName   string             `db:"-" json:"first_name" yaml:"first_name"`
	LastName    string             `db:"-" json:"last_name" yaml:"last_name"`
	IsStaff     bool               `db:"-" json:"is_staff" yaml:"is_staff"`
	IsSuperuser bool               `db:"-" json:"is_superuser" yaml:"is_superuser"`
	IsActive    bool               `db:"-" json:"is_active" yaml:"is_active"`
	IdentityID  string             `db:"-" json:"identity_id" yaml:"identity_id"`
	FacebookID  string             `db:"-" json:"facebook_id" yaml:"facebook_id"`
	LinkedinID  string             `db:"-" json:"linkedin_id" yaml:"linkedin_id"`
	GoogleID    string             `db:"-" json:"google_id" yaml:"google_id"`
	Phone       string             `db:"-" json:"phone" yaml:"phone"`
	IPAddress   string             `db:"-" json:"ip_address" yaml:"ip_address"`
	UserAgent   string             `db:"-" json:"user_agent" yaml:"user_agent"`
	Timezone    string             `db:"-" json:"timezone" yaml:"timezone"`
	Social      string             `db:"-" json:"social" yaml:"social"`
	Data        UserData           `db:"-" json:"data" yaml:"data"`
	LastLogin   pgtype.Timestamptz `db:"-" json:"last_login" yaml:"last_login"`
	CreatedAt   pgtype.Timestamptz `db:"-" json:"created_at" yaml:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"-" json:"updated_at" yaml:"updated_at"`
	ImageLinkID int                `db:"-" json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`

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
		"password":      model.Password,
		"email":         model.Email,
		"first_name":    model.FirstName,
		"last_name":     model.LastName,
		"is_staff":      model.IsStaff,
		"is_superuser":  model.IsSuperuser,
		"is_active":     model.IsActive,
		"identity_id":   model.IdentityID,
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
	return models.DefaultFields(&user)
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
