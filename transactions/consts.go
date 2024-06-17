package transactions

type (
	Status string
	Type   string
)

const (
	Wait      Status = "wait"
	Created   Status = "created"
	Pending   Status = "pending"
	Processed Status = "processed"
	Cancelled Status = "cancelled"
	Err       Status = "failed"

	Deposit    Type = "deposit"
	Withdrawal Type = "withdrawal"
	Invest     Type = "investment"
)
