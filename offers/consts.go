package offers

type (
	Type         string
	SecurityType string
	RegType      string
)

const (
	NewOffer            Type = "new"
	Draft               Type = "draft"
	LegalReview         Type = "legal-review"
	LegalDeclined       Type = "legal-declined"
	LegalAccepted       Type = "legal-accepted"
	Published           Type = "published"
	ClosedSuccessfully  Type = "closed-successfully"
	ClosedUnsuccesfully Type = "closed-unsuccesfully"

	Equity           SecurityType = "equity"
	PreferredEquity  SecurityType = "preferred-equity"
	Debt             SecurityType = "debt"
	ConvertibleDebt  SecurityType = "convertible-debt"
	EquityWarrants   SecurityType = "equity-warrants"
	ConvertibleBonds SecurityType = "convertible-bonds"
	PreferenceShares SecurityType = "preference-shares"

	REG_A       RegType = "reg A"
	REG_A_PLUG  RegType = "reg A+"
	REG_D_506_B RegType = "reg D 506(b)"
	REG_D_506_C RegType = "reg D 506(c)"
)
