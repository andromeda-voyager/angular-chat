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
	Permissions uint8 `json:"permissions"`
	ws          *websocket.Conn
}

var loggedInUsers map[string]*user.User
var connections map[int]map[int]*connection

func init() {
	loggedInUsers = make(map[string]*user.User)
	connections = make(map[int]map[int]*connection)
}

// ServerID => channels[], ws, serverPermissions

// Add .
func Add(user *user.User) *http.Cookie {
	token := random.NewSecureString(32)
	loggedInUsers[token] = user
	cookie := &http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}

// func AddConnection(serverID, accountID int, c *connection) {
// 	connections[serverID][accountID] = c
// }

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
func Post(a *user.Account, post server.Post) {
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
