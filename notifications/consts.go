package notifications

type (
	StatusType string
	Type       string
)

const (
	TypeDistribution Type = "distribution"
	TypeGeneral      Type = "general"
	TypeInternal     Type = "internal"
	TypeInvestment   Type = "investment"
	TypeOffer        Type = "offer"
	TypeProfile      Type = "profile"
	TypeSecurity     Type = "security"
	TypeSystem       Type = "system"
	TypeUser         Type = "user"
	TypeWallet       Type = "wallet"
	TypeFiler        Type = "filer"

	StatusRead   StatusType = "read"
	StatusUnread StatusType = "unread"
)
