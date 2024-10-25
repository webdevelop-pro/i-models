package offers

import (
	"github.com/webdevelop-pro/i-models/pgtype"
)

type OfferOffer struct {
	ID                int                `json:"id" yaml:"id"`
	UserID            int                `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Name              string             `json:"name" yaml:"name"`
	Slug              string             `json:"slug" yaml:"slug"`
	MinInvestment     int                `json:"min_investment" yaml:"min_investment"`
	Description       string             `json:"description" yaml:"description"`
	Title             string             `json:"title" yaml:"title"`
	Highlights        string             `json:"highlights" yaml:"highlights"`
	Valuation         float64            `json:"valuation" yaml:"valuation"`
	TotalShares       int                `json:"total_shares" yaml:"total_shares"`
	PricePerShare     float64            `json:"price_per_share" yaml:"price_per_share"`
	Status            NullOfferT         `json:"status,omitempty" yaml:"status,omitempty"`
	SecurityType      OfferSecurityTypeT `json:"security_type" yaml:"security_type"`
	Notes             string             `json:"notes" yaml:"notes"`
	SubscribedShares  int                `json:"subscribed_shares" yaml:"subscribed_shares"`
	ConfirmedShares   int                `json:"confirmed_shares" yaml:"confirmed_shares"`
	AdditionalDetails string             `json:"additional_details" yaml:"additional_details"`
	SeoTitle          string             `json:"seo_title" yaml:"seo_title"`
	SeoDescription    string             `json:"seo_description" yaml:"seo_description"`
	Website           string             `json:"website" yaml:"website"`
	State             string             `json:"state" yaml:"state"`
	City              string             `json:"city" yaml:"city"`
	Address1          string             `json:"address1" yaml:"address1"`
	Address2          string             `json:"address2" yaml:"address2"`
	ZipCode           string             `json:"zip_code" yaml:"zip_code"`
	Country           string             `json:"country" yaml:"country"`
	EntityID          string             `json:"entity_id" yaml:"entity_id"`
	Data              any                `json:"data" yaml:"data"`
	StartAt           pgtype.Timestamptz `json:"start_at,omitempty" yaml:"start_at,omitempty"`
	CloseAt           pgtype.Timestamptz `json:"close_at,omitempty" yaml:"close_at,omitempty"`
	ApprovedAt        pgtype.Timestamptz `json:"approved_at,omitempty" yaml:"approved_at,omitempty"`
	CreatedAt         pgtype.Timestamptz `json:"created_at" yaml:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at" yaml:"updated_at"`
	ImageLinkID       int                `json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`
	WalletID          int                `json:"wallet_id,omitempty" yaml:"wallet_id,omitempty"`
	LegalName         string             `json:"legal_name" yaml:"legal_name"`
	Twitter           string             `json:"twitter" yaml:"twitter"`
	Facebook          string             `json:"facebook" yaml:"facebook"`
	Github            string             `json:"github" yaml:"github"`
	Instagram         string             `json:"instagram" yaml:"instagram"`
	Telegram          string             `json:"telegram" yaml:"telegram"`
	Linkedin          string             `json:"linkedin" yaml:"linkedin"`
	Mastodon          string             `json:"mastodon" yaml:"mastodon"`
	EsignID           string             `json:"esign_id" yaml:"esign_id"`
	RegType           string             `json:"reg_type" yaml:"reg_type"`
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
		"wallet_id":          model.WalletID,
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
		"wallet_id",
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
	return "offer_offers"
}

func (model OfferOffer) GetID() any {
	return model.ID
}
