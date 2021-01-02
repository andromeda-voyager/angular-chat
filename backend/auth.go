package main

import (
	"bytes"
	"fmt"
	"nebula/message"
	"nebula/util"

	"golang.org/x/crypto/argon2"
)

var codes map[string]string

func init() {
	codes = make(map[string]string)
}

// IsCodeValid checks to see if code provided matches code sent by email
func isCodeValid(code, email string) bool {
	if codes[email] == code {
		delete(codes, email)
		return true
	}
	return false
}

// GenerateCode creates a code and stores it for later validation
func generateCode(email string) string {
	code := util.NewRandomString(5)
	codes[email] = code
	fmt.Println(code)
	return code
}

// SendCodeToEmail .
func SendCodeToEmail(email string) {
	if IsEmailInUse(email) {
		message.SendEmail([]byte("An account already exists with this email."), email)
	} else {
		msg := []byte("Nebula\n\nVerifcation Code:\t" + generateCode(email))
		fmt.Println(msg)
		message.SendEmail(msg, email)
	}
}

// Credentials holds the fields needed to authorize a user for login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(username, password string) {

}

// IsPasswordCorrect checks to see if the password in the database matches the password used to login
func IsPasswordCorrect(c Credentials) bool {
	storedPasswordHash, salt := getPassword(c.Email)
	passwordHashToVerify := argon2.IDKey([]byte(c.Password), []byte(salt), 4, 32*1024, 4, 32)
	return bytes.Compare(storedPasswordHash, passwordHashToVerify) == 0
}
