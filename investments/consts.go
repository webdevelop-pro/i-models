package investments

type (
	Status        string
	FundingStatus string
	FundingType   string
	EscrowType    string
	PaymentType   string
)

const (
	FundingCreationError FundingStatus = "creation_error"
	FundingNew           FundingStatus = "new"
	FundingInitialize    FundingStatus = "initialize"
	FundingInProgress    FundingStatus = "in_progress"
	FundingReceived      FundingStatus = "recieved" // ToDo: Send mistake here and in migrations
	FundingSettled       FundingStatus = "settled"
	FundingFailed        FundingStatus = "failed"
	FundingSentPending   FundingStatus = "sent_back_pending"
	FundingSentSettled   FundingStatus = "sent_back_settled"
	FundingCanceled      FundingStatus = "canceled"

	StatusNew              Status = "new"
	StatusConfirmed        Status = "confirmed"
	StatusLegallyConfirmed Status = "legally_confirmed"
	//nolint:misspell
	StatusClosedUnsuccesfully          Status = "closed_unsucessfully"
	StatusClosedSuccesfully            Status = "closed_successfully"
	StatusCancelledDuringInvestment    Status = "cancelled_during_investment"
	StatusCancelledAfterInvestment     Status = "cancelled_after_investment"
	StatusCancelledByManagerInvestment Status = "cancelled_by_manager"

	FundingNone   FundingType = "none"
	FundingWire   FundingType = "wire"
	FundingACH    FundingType = "ach"
	FundingWallet FundingType = "wallet"

	EscrowNone EscrowType = "none"
	EscrowNC   EscrowType = "north_capital"
	EscrowPT   EscrowType = "prime_trust"
	EscrowAG   EscrowType = "apex_group"

	PaymentNone PaymentType = "none"
	PaymentNC   PaymentType = "north_capital"
	PaymentPT   PaymentType = "prime_trust"
	PaymentAG   PaymentType = "apex_group"
)
