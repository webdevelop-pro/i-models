package fundingsource

type (
	Status string
	Type   string
)

const (
	StatusCreated  Status = "created"
	StatusVerified Status = "verified"
	StatusError    Status = "error"

	TypeBank     Type = "bank"
	TypeChecking Type = "checking"
	TypeSaving   Type = "savings"
)
