package main

import (
	"database/sql"
	"nebula/config"
	"nebula/util"
	"net/http"
	"time"
)

var loggedInUsers map[string]Account

func init() {
	loggedInUsers = make(map[string]Account)
}

// addSession .
func addSession(a *Account) http.Cookie {
	token := util.NewRandomSecureString(32)
	loggedInUsers[token] = *a
	cookie := http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}

// GetSession .
func GetSession(token string) Account {
	return loggedInUsers[token]
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
