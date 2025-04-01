package profiles

import (
	"reflect"
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
	ID                  int             `db:"-"`
	UserID              int             `db:"-"`
	Type                ProfileType     `db:"-"`
	Data                ProfileData     `db:"-"`
	KycID               *string         `db:"-"`
	KycStatus           KycType         `db:"-"`
	AccreditationID     *string         `db:"-"`
	AccreditationData   []Accreditation `db:"-"`
	AccreditationStatus *AccType        `db:"-"`
	EscrowID            string          `db:"-"`
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
	obj := &prof
	var res = []string{}
	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		dbtag := string(val.Type().Field(i).Tag.Get("db"))
		if dbtag != "" {
			res = append(res, dbtag)
		}
	}
	return res
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
	}
}

func (prof Profile) GetID() any {
	return prof.ID
}

func (prof *Profile) SetID(id any) {
	prof.ID = id.(int)
}
