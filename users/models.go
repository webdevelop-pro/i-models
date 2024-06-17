package users

// ToDo
// Add all fields
type UserData struct {
	NCIssuerID string `json:"nc_issuer_id"`
}

type User struct {
	ID         int      `json:"id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Email      string   `json:"email"`
	IdentityID string   `json:"identity_id"`
	Data       UserData `json:"data"`
}
