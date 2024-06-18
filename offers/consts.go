package offers

type (
	Type         string
	SecurityType string
	RegType      string
)

const (
	TypeNewOffer            Type = "new"
	TypeDraft               Type = "draft"
	TypeLegalReview         Type = "legal-review"
	TypeLegalDeclined       Type = "legal-declined"
	TypeLegalAccepted       Type = "legal-accepted"
	TypePublished           Type = "published"
	TypeClosedSuccessfully  Type = "closed-successfully"
	TypeClosedUnsuccesfully Type = "closed-unsuccesfully"

	SecurityEquity           SecurityType = "equity"
	SecurityPreferredEquity  SecurityType = "preferred-equity"
	SecurityDebt             SecurityType = "debt"
	SecurityConvertibleDebt  SecurityType = "convertible-debt"
	SecurityEquityWarrants   SecurityType = "equity-warrants"
	SecurityConvertibleBonds SecurityType = "convertible-bonds"
	SecurityPreferenceShares SecurityType = "preference-shares"

	RegA       RegType = "reg A"
	RegA_PLUG  RegType = "reg A+"
	RegD_506_B RegType = "reg D 506(b)"
	RegD_506_C RegType = "reg D 506(c)"
	RegCF      RegType = "reg CF"
)
