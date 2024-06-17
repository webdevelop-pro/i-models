package fundingsource

type (
	Status string
	Type   string
)

const (
	Created  Status = "created"
	Verified Status = "verified"
	Error    Status = "error"

	Bank     Type = "bank"
	Checking Type = "checking"
)
