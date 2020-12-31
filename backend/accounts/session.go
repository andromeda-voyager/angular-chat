package accounts

import (
	"database/sql"
	"fmt"
	"nebula/config"
	"nebula/message"
	"nebula/util"
	"net/http"
	"time"
)

var loggedInUsers map[string]string
var codes map[string]string

func init() {
	loggedInUsers = make(map[string]string)
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

// Add a user
func addSession(email string) http.Cookie {
	token := util.NewRandomSecureString(32)
	loggedInUsers[token] = email
	cookie := http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}

// IsEmailInUse checks if an email is already used for an account
func IsEmailInUse(email string) bool {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM Users WHERE Email=?;", email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	return rows.Next()
}
