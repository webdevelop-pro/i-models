package logs

type (
	ServiceType string
	Type        string
)

const (
	ServiceOry          ServiceType = "ory"
	ServicePlaid        ServiceType = "plaid"
	ServiceTwilio       ServiceType = "twilio"
	ServiceDwolla       ServiceType = "dwolla"
	ServiceSendGrid     ServiceType = "sendgrid"
	ServicePandaDoc     ServiceType = "pandadoc"
	ServiceHelloSign    ServiceType = "hellosign"
	ServiceNorthCapital ServiceType = "north_capital"

	TypeOutcoming Type = "outcoming"
	TypeIncoming  Type = "incoming"

	ObjectLogentry      = "logentry"
	ObjectPermission    = "permission"
	ObjectGroup         = "group"
	ObjectContentType   = "contenttype"
	ObjectSession       = "session"
	ObjectAccount       = "account"
	ObjectFiler         = "filer"
	ObjectOfferFiler    = "offerfiler"
	ObjectOffer         = "offer"
	ObjectProfile       = "profile"
	ObjectInvestment    = "investment"
	ObjectApplog        = "applog"
	ObjectEmail         = "email"
	ObjectComment       = "comment"
	ObjectWallet        = "wallet"
	ObjectFundingsource = "fundingsource"
	ObjectTransaction   = "transaction"
)
