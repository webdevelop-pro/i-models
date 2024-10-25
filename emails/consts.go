package emails

import "errors"

type EmailStatusT string

// Enum values for EmailStatusT
const (
	EmailStatusTDraft      EmailStatusT = "draft"
	EmailStatusTNew        EmailStatusT = "new"
	EmailStatusTInProgress EmailStatusT = "in_progress"
	EmailStatusTSent       EmailStatusT = "sent"
	EmailStatusTCancelled  EmailStatusT = "cancelled"
	EmailStatusTError      EmailStatusT = "error"
)

func AllEmailStatusT() []EmailStatusT {
	return []EmailStatusT{
		EmailStatusTDraft,
		EmailStatusTNew,
		EmailStatusTInProgress,
		EmailStatusTSent,
		EmailStatusTCancelled,
		EmailStatusTError,
	}
}

func (e EmailStatusT) IsValid() error {
	switch e {
	case EmailStatusTDraft, EmailStatusTNew, EmailStatusTInProgress, EmailStatusTSent, EmailStatusTCancelled, EmailStatusTError:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e EmailStatusT) String() string {
	return string(e)
}
