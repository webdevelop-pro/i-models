package filers

type OfferFileType string

const (
	OfferTypeCompany              OfferFileType = "company"
	OfferTypeTax                  OfferFileType = "tax"
	OfferTypeInvestmentAgreements OfferFileType = "investment-agreements"
	OfferTypeInvestorUpdates      OfferFileType = "investor-updates"
	OfferTypeOther                OfferFileType = "other"
)
