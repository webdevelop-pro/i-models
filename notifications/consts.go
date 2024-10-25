package notifications

import "github.com/friendsofgo/errors"

type NotificationStatusT string

// Enum values for NotificationStatusT
const (
	NotificationStatusTUnread NotificationStatusT = "unread"
	NotificationStatusTRead   NotificationStatusT = "read"
)

func AllNotificationStatusT() []NotificationStatusT {
	return []NotificationStatusT{
		NotificationStatusTUnread,
		NotificationStatusTRead,
	}
}

func (e NotificationStatusT) IsValid() error {
	switch e {
	case NotificationStatusTUnread, NotificationStatusTRead:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e NotificationStatusT) String() string {
	return string(e)
}

type NotificationTypeT string

// Enum values for NotificationTypeT
const (
	NotificationTypeTGeneral      NotificationTypeT = "general"
	NotificationTypeTInvestment   NotificationTypeT = "investment"
	NotificationTypeTOffer        NotificationTypeT = "offer"
	NotificationTypeTProfile      NotificationTypeT = "profile"
	NotificationTypeTWallet       NotificationTypeT = "wallet"
	NotificationTypeTUser         NotificationTypeT = "user"
	NotificationTypeTFiler        NotificationTypeT = "filer"
	NotificationTypeTDistribution NotificationTypeT = "distribution"
	NotificationTypeTSecurity     NotificationTypeT = "security"
	NotificationTypeTSystem       NotificationTypeT = "system"
	NotificationTypeTInternal     NotificationTypeT = "internal"
)

func AllNotificationTypeT() []NotificationTypeT {
	return []NotificationTypeT{
		NotificationTypeTGeneral,
		NotificationTypeTInvestment,
		NotificationTypeTOffer,
		NotificationTypeTProfile,
		NotificationTypeTWallet,
		NotificationTypeTUser,
		NotificationTypeTFiler,
		NotificationTypeTDistribution,
		NotificationTypeTSecurity,
		NotificationTypeTSystem,
		NotificationTypeTInternal,
	}
}

func (e NotificationTypeT) IsValid() error {
	switch e {
	case NotificationTypeTGeneral, NotificationTypeTInvestment, NotificationTypeTOffer, NotificationTypeTProfile, NotificationTypeTWallet, NotificationTypeTUser, NotificationTypeTFiler, NotificationTypeTDistribution, NotificationTypeTSecurity, NotificationTypeTSystem, NotificationTypeTInternal:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e NotificationTypeT) String() string {
	return string(e)
}
