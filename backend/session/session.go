package session

import (
	"nebula/server"
	"nebula/user"
	"nebula/util"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Connection struct {
	Permissions uint8 `json:"permissions"`
	ws          *websocket.Conn
}

var loggedInUsers map[string]*user.User
var connections map[int]map[int]*Connection

func init() {
	loggedInUsers = make(map[string]*user.User)
	connections = make(map[int]map[int]*Connection)
}

// ServerID => channels[], ws, serverPermissions

// Add .
func Add(user *user.User) *http.Cookie {
	token := util.NewRandomSecureString(32)
	loggedInUsers[token] = user
	cookie := &http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}

func AddConnection(serverID, accountID int, c *Connection) {
	connections[serverID][accountID] = c
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
func Post(a *user.Account, post server.Post) {
	// post.TimePosted = time.Now().UTC()
	// post.AccountID = a.ID
	// connectionsToServer := serverConnections[post.ServerID]
	// for _, connection := range connectionsToServer {
	// 	connection.Send(post)
	// }
}
