package wallets

type (
	Status string
)

const (
	StatusCreated  Status = "created"
	StatusVerified Status = "verified"
	StatusError    Status = "error"
)
