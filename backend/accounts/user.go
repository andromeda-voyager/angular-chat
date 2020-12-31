package accounts

// User holds all the values stored for each user in the database. Used to unmarshal json sent during account creation.
type User struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	AvatarURL string   `json:"avatarURL"`
	Code      string   `json:"code"`
	Servers   []string `json:"servers"`
}

func (user *User) hasValidFields() bool {
	if len(user.Email) < 5 {
		return false
	}
	if len(user.Username) < 1 {
		return false
	}
	if len(user.Password) < 8 {
		return false
	}
	return true
}
