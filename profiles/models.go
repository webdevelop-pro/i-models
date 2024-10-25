package profiles

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/models"
)

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
	store               *models.Store[Profile]
	ID                  int             `db:"id"`
	UserID              int             `db:"user_id"`
	Type                ProfileType     `db:"type"`
	Data                ProfileData     `db:"data"`
	KycID               *string         `db:"kyc_id"`
	KycStatus           KycType         `db:"kyc_status"`
	AccreditationID     *string         `db:"accreditation_id"`
	AccreditationData   []Accreditation `db:"accreditation_data"`
	AccreditationStatus AccType         `db:"accreditation_status"`
	EscrowID            string          `db:"escrow_id"`
	WalletID            *int            `db:"wallet_id"`
}

func New(pool *db.DB) *Profile {
	profile := Profile{}
	store := models.NewStore[Profile](pool, profile.Table(), profile)
	profile.store = store
	return &profile
}

func (prof *Profile) SelectByID(ctx context.Context, id int) (*Profile, error) {
	profile, err := prof.store.Select(ctx, sq.Eq{"id": id})
	profile.store = prof.store
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (prof *Profile) Select(ctx context.Context, params map[string]any) (*Profile, error) {
	profile, err := prof.store.Select(ctx, params)
	profile.store = prof.store
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (prof Profile) Table() string {
	return "investment_profiles"
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

func (prof Profile) Get(key string) (any, error) {
	switch key {
	case "kyc_id":
		return prof.KycID, nil
	case "kyc_status":
		return prof.KycStatus, nil
	default:
		return nil, errors.Errorf("unknown key %s", key)
	}
}

func (prof *Profile) SetUserID(userID int) {
	prof.UserID = userID
}

func (prof *Profile) GetUserID() int {
	return prof.UserID
}

func (prof *Profile) SetKycID(kycID string) {
	prof.KycID = &kycID
}

func (prof *Profile) GetKycID() *string {
	return prof.KycID
}

func (prof *Profile) SetKycStatus(kycStatus KycType) {
	prof.KycStatus = kycStatus
}

func (prof *Profile) ToDTO() *Profile {
	return nil
	/*
		return &Profile{
			ID:                  prof.ID,
			UserID:              prof.UserID,
			Type:                prof.Type,
			Data:                prof.Data,
			KycID:               prof.KycID,
			KycStatus:           prof.KycStatus,
			AccreditationID:     prof.AccreditationID,
			AccreditationData:   prof.AccreditationData,
			AccreditationStatus: prof.AccreditationStatus,
			EscrowID:            prof.EscrowID,
			WalletID:            prof.WalletID,
		}
	*/
}
