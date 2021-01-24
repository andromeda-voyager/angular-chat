package user

import (
	"fmt"
	"nebula/random"
)

// Credentials holds the fields needed to authorize a user for login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var codes map[string]string

func init() {
	codes = make(map[string]string)
}

// // IsPasswordCorrect checks to see if the password in the database matches the password used to login
// func isPasswordCorrect(salt, hashedPassword []byte, passwordAttempt string) bool {
// 	//storedPasswordHash, salt := getPassword(c.Email)
// 	hash := argon2.IDKey([]byte(passwordAttempt), []byte(salt), 4, 32*1024, 4, 32)
// 	return bytes.Compare(hashedPassword, hash) == 0
// }

// IsCodeValid checks to see if code provided matches code sent by email
func IsCodeValid(code, email string) bool {
	if codes[email] == code {
		delete(codes, email)
		return true
	}
	return false
}

// GenerateCode creates a code and stores it for later validation
func generateCode(email string) string {
	code := random.NewSecureString(5)
	codes[email] = code
	fmt.Println(code)
	return code
}
