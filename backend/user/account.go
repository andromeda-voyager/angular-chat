package user

import (
	"database/sql"
	"nebula/config"
)

// Account holds the json values sent by a client to create a new account
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

func (a *Account) hasValidFields() bool {
	if len(a.Email) < 5 {
		return false
	}
	if len(a.Username) < 1 {
		return false
	}
	if len(a.Password) < 8 {
		return false
	}
	return true
}

// IsEmailInUse checks if an email is already used for an account
func IsEmailInUse(email string) bool {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM Account WHERE email=?;", email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	return rows.Next()
}
