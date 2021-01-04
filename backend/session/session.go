package session

import (
	"nebula/account"
	"nebula/server"
	"nebula/util"
	"net/http"
	"time"
)

var loggedInUsers map[string]*account.Account
var serverConnections map[string]map[string]server.Connection

func init() {
	loggedInUsers = make(map[string]*account.Account)
	serverConnections = make(map[string]map[string]server.Connection)
}

// Add .
func Add(a *account.Account) http.Cookie {
	token := util.NewRandomSecureString(32)
	loggedInUsers[token] = a
	cookie := http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}

// Get .
func Get(token string) *account.Account {
	return loggedInUsers[token]
}

// Post .
func Post(a *account.Account, post server.Post) {
	// send post to open connections
}
