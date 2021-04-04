package main

import (
	"nebula/server"
	"nebula/user"
	"net/http"

	router "github.com/andromeda-voyager/go-router"
	"golang.org/x/net/websocket"
)

// LoginResponse .
type LoginResponse struct {
	User    *user.User       `json:"user"`
	Servers []*server.Server `json:"servers"`
}

// ServerRequest .
type ServerRequest struct {
	ServerID  int            `json:"serverID"`
	ChannelID int            `json:"channelID"`
	Channel   server.Channel `json:"channel"`
	Role      server.Role    `json:"role"`
	Roles     []server.Role  `json:"roles"`
}

func keepConnectionOpen(ws *websocket.Conn) {
	for {
		msg := make([]byte, 0)
		_, err := ws.Read(msg)
		if err != nil {
			defer ws.Close()
			return
		}
	}
}

func socket(ws *websocket.Conn) {
	cookie, err := ws.Request().Cookie("Auth")
	if err != nil {
		defer ws.Close()
	} else {
		user, _ := user.GetSession((cookie.Value))
		server.AddConnection(user.ID, ws)
		keepConnectionOpen(ws)
		// if err := websocket.JSON.Send(ws, loginResponse); err != nil {
		// 	fmt.Println("unable to send")
		// 	defer ws.Close()
		// }
	}

}

func init() {

	authGroup := router.NewGroup()
	authGroup.Use(user.Authenticate)

	authGroup.Get("/ws", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		s := websocket.Handler(socket)
		s.ServeHTTP(w, r)
	})

	authGroup.Get("/test-websocket", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		server.SendMessage(u.ID, "Lo")
	})

}
