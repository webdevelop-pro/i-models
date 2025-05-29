package logs

import "github.com/pkg/errors"

type (
	ServiceType string
	ObjectType  string
	Type        string
)

const (
	ObjectLogentry      ObjectType = "logentry"
	ObjectPermission    ObjectType = "permission"
	ObjectGroup         ObjectType = "group"
	ObjectContentType   ObjectType = "contenttype"
	ObjectSession       ObjectType = "session"
	ObjectAccount       ObjectType = "account"
	ObjectFiler         ObjectType = "filer"
	ObjectOfferFiler    ObjectType = "offerfiler"
	ObjectOffer         ObjectType = "offer"
	ObjectProfile       ObjectType = "profile"
	ObjectInvestment    ObjectType = "investment"
	ObjectApplog        ObjectType = "applog"
	ObjectEmail         ObjectType = "email"
	ObjectComment       ObjectType = "comment"
	ObjectWallet        ObjectType = "wallet"
	ObjectFundingsource ObjectType = "fundingsource"
	ObjectTransaction   ObjectType = "transaction"

	MsgContentIDNotFound = "content id not found"
)

type ServicesT string

// Enum values for ServicesT
const (
	ServicesTNorthCapital ServicesT = "north_capital"
	ServicesTPlaid        ServicesT = "plaid"
	ServicesTSendgrid     ServicesT = "sendgrid"
	ServicesTHellosign    ServicesT = "hellosign"
	ServicesTPandadoc     ServicesT = "pandadoc"
	ServicesTTwilio       ServicesT = "twilio"
	ServicesTOry          ServicesT = "ory"
	ServicesTDwolla       ServicesT = "dwolla"
)

func AllServicesT() []ServicesT {
	return []ServicesT{
		ServicesTNorthCapital,
		ServicesTPlaid,
		ServicesTSendgrid,
		ServicesTHellosign,
		ServicesTPandadoc,
		ServicesTTwilio,
		ServicesTOry,
		ServicesTDwolla,
	}
}

func (e ServicesT) IsValid() error {
	switch e {
	case ServicesTNorthCapital, ServicesTPlaid, ServicesTSendgrid, ServicesTHellosign, ServicesTPandadoc, ServicesTTwilio, ServicesTOry, ServicesTDwolla:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e ServicesT) String() string {
	return string(e)
}

type LogTypeT string

// Enum values for LogTypeT
const (
	LogTypeTOutcoming LogTypeT = "outcoming"
	LogTypeTIncoming  LogTypeT = "incoming"
)

func AllLogTypeT() []LogTypeT {
	return []LogTypeT{
		LogTypeTOutcoming,
		LogTypeTIncoming,
	}
}

func (e LogTypeT) IsValid() error {
	switch e {
	case LogTypeTOutcoming, LogTypeTIncoming:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e LogTypeT) String() string {
	return string(e)
}
