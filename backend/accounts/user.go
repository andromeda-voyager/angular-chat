package accounts

// User holds all the values stored for each user in the database. Used to unmarshal json sent during account creation.
type User struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarURL"`
	Code      string `json:"code"`
}
