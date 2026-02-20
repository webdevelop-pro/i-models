package filers

type OfferFileType string

const (
	OfferTypeCompany              OfferFileType = "company"
	OfferTypeTax                  OfferFileType = "tax"
	OfferTypeInvestmentAgreements OfferFileType = "investment-agreements"
	OfferTypeInvestorUpdates      OfferFileType = "investor-updates"
	OfferTypeOther                OfferFileType = "other"

	AppLabel  = "filer_ltree"
	ModelName = "filerltree"
	TableName = "filer_filers_ltree"

	pkgName = "models/filerltree"
)
