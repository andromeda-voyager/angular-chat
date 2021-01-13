package server

import (
	"golang.org/x/net/websocket"
)

// Connection .
type Connection struct {
	Server      *Server `json:"server"`
	Alias       string  `json:"alias"`
	Permissions uint8   `json:"permissions"`
	ws          *websocket.Conn
}

// Send .
func (c *Connection) Send(post Post) {
	websocket.JSON.Send(c.ws, post)
}
