package profiles

import (
	"github.com/webdevelop-pro/go-common/db"
)

const Table = "investment_profiles"

type ProfileData struct {
	Citizenship CitizenType `json:"citizenship"`
	FirstName   string      `json:"first_name"`
	MiddleName  string      `json:"middle_name"`
	LastName    string      `json:"last_name"`
	DOB         string      `json:"dob"`
	Country     string      `json:"country"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Phone       string      `json:"phone"`
	Address1    string      `json:"address1"`
	Address2    string      `json:"address2"`
	ZipCode     string      `json:"zip_code"`
	SSN         string      `json:"ssn"`
	IPAddress   string      `json:"ip_address"`
	NCPartyID   string      `json:"nc_party_id"`
	NCLinkID    string      `json:"nc_link_id"`
}

type Profile struct {
	ID                  int             `db:"id"`
	UserID              int             `db:"user_id"`
	Type                ProfileType     `db:"type"`
	Data                ProfileData     `db:"data"`
	KycID               *string         `db:"kyc_id"`
	KycStatus           KycType         `db:"kyc_status"`
	AccreditationID     *string         `db:"accreditation_id"`
	AccreditationData   []Accreditation `db:"accreditation_data"`
	AccreditationStatus *AccType        `db:"accreditation_status"`
	EscrowID            string          `db:"escrow_id"`
	WalletID            *int            `db:"wallet_id"`
}

func New(pool *db.DB) *Profile {
	profile := Profile{}
	return &profile
}

func (prof Profile) Table() string {
	return Table
}

func (prof Profile) PrimaryFieldKey() string {
	return "id"
}

func (prof Profile) PrimaryFieldValue() any {
	return prof.ID
}

func (prof Profile) Fields() []string {
	return []string{""}
}

func (prof Profile) ToJSON() map[string]any {
	return map[string]any{
		"ID":                  prof.ID,
		"UserID":              prof.UserID,
		"Type":                prof.Type,
		"Data":                prof.Data,
		"KycID":               prof.KycID,
		"KycStatus":           prof.KycStatus,
		"AccreditationID":     prof.AccreditationID,
		"AccreditationData":   prof.AccreditationData,
		"AccreditationStatus": prof.AccreditationStatus,
		"EscrowID":            prof.EscrowID,
		"WalletID":            prof.WalletID,
	}
}

func (prof Profile) GetID() any {
	return prof.ID
}
