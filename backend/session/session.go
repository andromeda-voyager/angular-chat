package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"nebula/random"
	"nebula/server"
	"nebula/user"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

//map(key server id) of map(key userID) of *connection
var servers map[int]map[int]*connection

//map(key channeld id) of map(key userID) of *connection
var channels map[int]map[int]*connection

var loggedInUsers map[string]*user.User
var connections map[int]*connection

type connection struct {
	ws        *websocket.Conn
	userID    int
	serverID  int
	channelID int
}

func init() {
	loggedInUsers = make(map[string]*user.User)
	servers = make(map[int]map[int]*connection)
	channels = make(map[int]map[int]*connection)
	connections = make(map[int]*connection)
}

func ConnectToChannel(channelID, userID int) {
	c, ok := connections[userID]
	if ok {
		if c.serverID != 0 {
			if c.channelID != 0 {
				disconnectFromChannel(c.channelID, userID)
			}
			c.channelID = channelID
			_, ok := channels[channelID]
			if !ok {
				channels[channelID] = make(map[int]*connection)
			}
			channels[channelID][userID] = c
		}
	}

}

func ConnectToServer(serverID, userID int) {
	c, ok := connections[userID]
	if ok {
		if c.serverID != 0 {
			if c.channelID != 0 {
				disconnectFromChannel(c.channelID, userID)
			}
			disconnectFromServer(c.serverID, c.channelID, userID)
		}
		_, ok := servers[c.serverID]
		if !ok {
			servers[c.serverID] = make(map[int]*connection)
		}
		servers[c.serverID][userID] = c
	}
}

func disconnectFromServer(serverID, channelID, userID int) {
	server, ok := servers[serverID]
	if ok {
		_, ok := server[userID]
		if ok {
			delete(server, userID)
		}
	}
}

func disconnectFromChannel(channelID, userID int) {
	channel, ok := channels[channelID]
	if ok {
		_, ok := channel[userID]
		if ok {
			delete(channel, userID)
		}
	}
}

func AddConnection(userID int, ws *websocket.Conn) {
	connections[userID] = &connection{userID: userID, ws: ws, serverID: 0, channelID: 0}
}

func SendMessage(userID int, message string) {
	c, ok := connections[userID]
	if ok {
		if err := websocket.JSON.Send(c.ws, "LO"); err != nil {
			fmt.Println("ws message sent to client")
		}
	}
}

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
func Get(token string) (*user.User, bool) {
	user, ok := loggedInUsers[token]
	return user, ok
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
