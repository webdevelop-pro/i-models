package transactions

type (
	Status string
	Type   string
)

const (
	StatusWait      Status = "wait"
	StatusCreated   Status = "created"
	StatusPending   Status = "pending"
	StatusProcessed Status = "processed"
	StatusCancelled Status = "cancelled"
	StatusErr       Status = "failed"

	TypeDeposit    Type = "deposit"
	TypeWithdrawal Type = "withdrawal"
	TypeInvest     Type = "investment"
)
