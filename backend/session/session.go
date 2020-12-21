package session

import (
	"nebula/util"
	"net/http"
	"time"
)

var loggedInUsers map[string]string

func init() {
	loggedInUsers = make(map[string]string)
}

// Add a user
func Add(email string) http.Cookie {
	token := util.NewRandomSecureString(32)
	loggedInUsers[token] = email
	cookie := http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}
