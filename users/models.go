package users

import (
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
	ID          int64              `json:"id" yaml:"id"`
	Password    string             `json:"password" yaml:"password"`
	Email       string             `json:"email" yaml:"email"`
	FirstName   string             `json:"first_name" yaml:"first_name"`
	LastName    string             `json:"last_name" yaml:"last_name"`
	IsStaff     bool               `json:"is_staff" yaml:"is_staff"`
	IsSuperuser bool               `json:"is_superuser" yaml:"is_superuser"`
	IsActive    bool               `json:"is_active" yaml:"is_active"`
	IdentityID  string             `json:"identity_id" yaml:"identity_id"`
	FacebookID  string             `json:"facebook_id" yaml:"facebook_id"`
	LinkedinID  string             `json:"linkedin_id" yaml:"linkedin_id"`
	GoogleID    string             `json:"google_id" yaml:"google_id"`
	Phone       string             `json:"phone" yaml:"phone"`
	IPAddress   string             `json:"ip_address" yaml:"ip_address"`
	UserAgent   string             `json:"user_agent" yaml:"user_agent"`
	Timezone    string             `json:"timezone" yaml:"timezone"`
	Social      string             `json:"social" yaml:"social"`
	Data        UserData           `json:"data" yaml:"data"`
	LastLogin   pgtype.Timestamptz `json:"last_login" yaml:"last_login"`
	CreatedAt   pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
	ImageLinkID int                `json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`
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

func (model UserUser) Fields() []string {
	return []string{
		"ID",
		"Password",
		"Email",
		"FirstName",
		"LastName",
		"IsStaff",
		"IsSuperuser",
		"IsActive",
		"IdentityID",
		"FacebookID",
		"LinkedinID",
		"GoogleID",
		"Phone",
		"IPAddress",
		"UserAgent",
		"Timezone",
		"Social",
		"Data",
		"LastLogin",
		"CreatedAt",
		"UpdatedAt",
		"ImageLinkID",
	}
}

func (model UserUser) Table() string {
	return "user_users"
}

func (model UserUser) GetID() any {
	return model.ID
}
