package session

import (
	"nebulous/util"
	"net/http"
	"time"
)

var loggedInUsers map[string]string

func init() {
	loggedInUsers = make(map[string]string)
}

func Add(email string) http.Cookie {
	token := util.NewRandomString(32)
	loggedInUsers[token] = email
	cookie := http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}
