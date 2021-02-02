package session

import (
	"crypto/rand"
	"encoding/base64"
	"nebula/random"
	"nebula/server"
	"nebula/user"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type connection struct {
	channels map[int]bool
	ws       *websocket.Conn
}

var loggedInUsers map[string]*user.User
var connections map[int]map[int]*connection

func init() {
	loggedInUsers = make(map[string]*user.User)
	connections = make(map[int]map[int]*connection)
}

// channelID -> map of tokens -> ws

//serverID -> map of tokens (for eached user logged into server)-> map of channels

// Add .
func Add(user *user.User) *http.Cookie {
	token := random.NewSecureString(32)
	loggedInUsers[token] = user
	cookie := &http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}

// Remove .
func Remove(token string) {
	_, ok := loggedInUsers[token]
	if ok {
		delete(loggedInUsers, token)
	}
}

// Get .
func Get(token string) *user.User {
	return loggedInUsers[token]
}

// Post .
func Post(a *user.Account, m server.Message) {
	// post.TimePosted = time.Now().UTC()
	// post.AccountID = a.ID
	// connectionsToServer := serverConnections[post.ServerID]
	// for _, connection := range connectionsToServer {
	// 	connection.Send(post)
	// }
}

// GetRandomBytes Returns a random slice of bytes
func GetRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	for err != nil {
		_, err = rand.Read(b)
	}
	return b
}

// NewRandomSecureString Returns a random cryptographically secure string
func NewRandomSecureString(n int) string {
	b := GetRandomBytes(n)
	return base64.RawStdEncoding.EncodeToString(b)
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
