package profiles

type (
	ProfileType string
	CitizenType string
	KycType     string
	AccType     string
)

const (
	AccNew          AccType = "new"
	AccInProgress   AccType = "in_progress"
	AccInfoRequired AccType = "info_required"
	AccApproved     AccType = "approved"
	AccDeclined     AccType = "declined"
	AccExpired      AccType = "expired"
)

const (
	KycNew      KycType = "new"
	KycPending  KycType = "pending"
	KycDeclined KycType = "declined"
	KycApproved KycType = "approved"
)

const (
	USACitizen  CitizenType = "U.S. Citizen"
	USAResident CitizenType = "U.S. Resident"
	NonResident CitizenType = "Non Resident"
)
