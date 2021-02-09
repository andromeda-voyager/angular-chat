package user

// Credentials holds the fields needed to authorize a user for login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
