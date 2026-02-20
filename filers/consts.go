package filers

type FolderName string
type Type string

const (
	FolderCompany              FolderName = "company"
	FolderTax                  FolderName = "tax"
	FolderInvestmentAgreements FolderName = "investment_agreements"
	FolderInvestorUpdates      FolderName = "investor_updates"
	FolderOther                FolderName = "other"

	TypeFile   Type = "file"
	TypeFolder Type = "folder"

	AppLabel  = "filer_ltree"
	ModelName = "filerltree"
	TableName = "filer_filers_ltree"

	pkgName = "models/filerltree"
)
