package server

import "golang.org/x/net/websocket"

// Connection .
type Connection struct {
	ServerID    int
	UserID      int
	Alias       string
	Permissions uint8
	Connection  *websocket.Conn
}
