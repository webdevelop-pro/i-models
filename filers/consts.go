package filers

type FolderName string
type Type string

const (
	FolderCompany              FolderName = "company"
	FolderTax                  FolderName = "tax"
	FolderInvestmentAgreements FolderName = "investment_agreements"
	FolderInvestorUpdates      FolderName = "investor_updates"
	FolderOther                FolderName = "other"

	TypeFile          Type = "file"
	TypeFolder        Type = "folder"
	TypeLink          Type = "link"
	TypeFileThumbnail Type = "file_thumbnail"

	AppLabel  = "filer_ltree"
	ModelName = "filerltree"
	TableName = "filer_filers_ltree"

	pkgName = "models/filerltree"
)

func AllType() []Type {
	return []Type{TypeFile, TypeFolder, TypeLink, TypeFileThumbnail}
}
