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
)
