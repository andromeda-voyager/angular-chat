package accounts

import (
	"bytes"

	"golang.org/x/crypto/argon2"
)

// Credentials holds the fields needed to authorize a user for login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(username, password string) {

}

func IsPasswordCorrect(c Credentials) bool {
	storedPasswordHash, salt := getPassword(c.Email)
	passwordHashToVerify := argon2.IDKey([]byte(c.Password), []byte(salt), 4, 32*1024, 4, 32)
	return bytes.Compare(storedPasswordHash, passwordHashToVerify) == 0
}
