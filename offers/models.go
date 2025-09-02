package offers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/i-models/models"
	"github.com/webdevelop-pro/i-models/pgtype"
)

type OfferOffer struct {
	ID                int                `db:"-" json:"id" yaml:"id"`
	UserID            int                `db:"-" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Name              string             `db:"-" json:"name" yaml:"name"`
	Slug              string             `db:"-" json:"slug" yaml:"slug"`
	MinInvestment     int                `db:"-" json:"min_investment" yaml:"min_investment"`
	Description       string             `db:"-" json:"description" yaml:"description"`
	Title             string             `db:"-" json:"title" yaml:"title"`
	Highlights        string             `db:"-" json:"highlights" yaml:"highlights"`
	Valuation         float64            `db:"-" json:"valuation" yaml:"valuation"`
	TotalShares       int                `db:"-" json:"total_shares" yaml:"total_shares"`
	PricePerShare     float64            `db:"-" json:"price_per_share" yaml:"price_per_share"`
	Status            NullOfferT         `db:"-" json:"status,omitempty" yaml:"status,omitempty"`
	SecurityType      OfferSecurityTypeT `db:"-" json:"security_type" yaml:"security_type"`
	Notes             string             `db:"-" json:"notes" yaml:"notes"`
	SubscribedShares  int                `db:"-" json:"subscribed_shares" yaml:"subscribed_shares"`
	ConfirmedShares   int                `db:"-" json:"confirmed_shares" yaml:"confirmed_shares"`
	AdditionalDetails string             `db:"-" json:"additional_details" yaml:"additional_details"`
	SeoTitle          string             `db:"-" json:"seo_title" yaml:"seo_title"`
	SeoDescription    string             `db:"-" json:"seo_description" yaml:"seo_description"`
	Website           string             `db:"-" json:"website" yaml:"website"`
	State             string             `db:"-" json:"state" yaml:"state"`
	City              string             `db:"-" json:"city" yaml:"city"`
	Address1          string             `db:"-" json:"address1" yaml:"address1"`
	Address2          string             `db:"-" json:"address2" yaml:"address2"`
	ZipCode           string             `db:"-" json:"zip_code" yaml:"zip_code"`
	Country           string             `db:"-" json:"country" yaml:"country"`
	EntityID          string             `db:"-" json:"entity_id" yaml:"entity_id"`
	Data              any                `db:"-" json:"data" yaml:"data"`
	StartAt           pgtype.Timestamptz `db:"-" json:"start_at,omitempty" yaml:"start_at,omitempty"`
	CloseAt           pgtype.Timestamptz `db:"-" json:"close_at,omitempty" yaml:"close_at,omitempty"`
	ApprovedAt        pgtype.Timestamptz `db:"-" json:"approved_at,omitempty" yaml:"approved_at,omitempty"`
	CreatedAt         pgtype.Timestamptz `db:"-" json:"created_at" yaml:"created_at"`
	UpdatedAt         pgtype.Timestamptz `db:"-" json:"updated_at" yaml:"updated_at"`
	ImageLinkID       int                `db:"-" json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`
	LegalName         string             `db:"-" json:"legal_name" yaml:"legal_name"`
	Twitter           string             `db:"-" json:"twitter" yaml:"twitter"`
	Facebook          string             `db:"-" json:"facebook" yaml:"facebook"`
	Github            string             `db:"-" json:"github" yaml:"github"`
	Instagram         string             `db:"-" json:"instagram" yaml:"instagram"`
	Telegram          string             `db:"-" json:"telegram" yaml:"telegram"`
	Linkedin          string             `db:"-" json:"linkedin" yaml:"linkedin"`
	Mastodon          string             `db:"-" json:"mastodon" yaml:"mastodon"`
	EsignID           string             `db:"-" json:"esign_id" yaml:"esign_id"`
	RegType           string             `db:"-" json:"reg_type" yaml:"reg_type"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func New(db db.Repository) *OfferOffer {
	return &OfferOffer{
		db:  db,
		fns: map[string]any{},
	}
}

func (model OfferOffer) ToJSON() map[string]any {
	return map[string]any{
		"id":                 model.ID,
		"user_id":            model.UserID,
		"name":               model.Name,
		"slug":               model.Slug,
		"min_investment":     model.MinInvestment,
		"description":        model.Description,
		"title":              model.Title,
		"highlights":         model.Highlights,
		"valuation":          model.Valuation,
		"total_shares":       model.TotalShares,
		"price_per_share":    model.PricePerShare,
		"status":             model.Status,
		"security_type":      model.SecurityType,
		"notes":              model.Notes,
		"subscribed_shares":  model.SubscribedShares,
		"confirmed_shares":   model.ConfirmedShares,
		"additional_details": model.AdditionalDetails,
		"seo_title":          model.SeoTitle,
		"seo_description":    model.SeoDescription,
		"website":            model.Website,
		"state":              model.State,
		"city":               model.City,
		"address1":           model.Address1,
		"address2":           model.Address2,
		"zip_code":           model.ZipCode,
		"country":            model.Country,
		"entity_id":          model.EntityID,
		"data":               model.Data,
		"start_at":           model.StartAt,
		"close_at":           model.CloseAt,
		"approved_at":        model.ApprovedAt,
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
		"image_link_id":      model.ImageLinkID,
		"legal_name":         model.LegalName,
		"twitter":            model.Twitter,
		"facebook":           model.Facebook,
		"github":             model.Github,
		"instagram":          model.Instagram,
		"telegram":           model.Telegram,
		"linkedin":           model.Linkedin,
		"mastodon":           model.Mastodon,
		"esign_id":           model.EsignID,
		"reg_type":           model.RegType,
	}
}

func (model OfferOffer) Fields() []string {
	return []string{
		"id",
		"user_id",
		"name",
		"slug",
		"min_investment",
		"description",
		"title",
		"highlights",
		"valuation",
		"total_shares",
		"price_per_share",
		"status",
		"security_type",
		"notes",
		"subscribed_shares",
		"confirmed_shares",
		"additional_details",
		"seo_title",
		"seo_description",
		"website",
		"state",
		"city",
		"address1",
		"address2",
		"zip_code",
		"country",
		"entity_id",
		"data",
		"start_at",
		"close_at",
		"approved_at",
		"created_at",
		"updated_at",
		"image_link_id",
		"legal_name",
		"twitter",
		"facebook",
		"github",
		"instagram",
		"telegram",
		"linkedin",
		"mastodon",
		"esign_id",
		"reg_type",
	}
}

func (model OfferOffer) Table() string {
	return TableName
}

func (model OfferOffer) GetID() any {
	return model.ID
}

func (model *OfferOffer) SetID(id any) {
	model.ID = id.(int)
}

func (model *OfferOffer) SetDB(db db.Repository) {
	model.db = db
}

func Get(ctx context.Context, db db.Repository, where map[string]any) (*OfferOffer, error) {
	model, err := models.RetriveOne[OfferOffer](
		ctx,
		db,
		where,
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot get model")
		return nil, err
	}
	model.db = db
	return model, nil
}
