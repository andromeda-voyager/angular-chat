package user

import (
	"fmt"
	"nebula/database"
	"nebula/util"
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

func getPassword(email string) ([]byte, []byte) {
	var args []interface{}
	args = append(args, email)
	rows, err := database.Query("SELECT password, salt FROM account WHERE email=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var password []byte
		var salt = make([]byte, 32)
		rows.Scan(&password, &salt)
		return password, salt
	}
	return nil, nil
}

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
	code := util.NewRandomString(5)
	codes[email] = code
	fmt.Println(code)
	return code
}
