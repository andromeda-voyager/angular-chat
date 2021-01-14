package server

import (
	"encoding/json"
	"fmt"
	"nebula/database"
	"nebula/util"
	"net/http"

	"golang.org/x/net/websocket"
)

// Server .
type Server struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
	Posts    []Post `json:"posts"`
}

// Invite .
type Invite struct {
	Code     string `json:"code"`
	ServerID string `json:"serverID"`
}

// New .
func New(r *http.Request) Server {
	serverStr := []byte(r.FormValue("server"))
	var server Server
	json.Unmarshal(serverStr, &server)
	server.ImageURL = util.SaveImage(r)
	var args []interface{}
	args = append(args, server.Name, server.ImageURL)
	fmt.Println(server.Name)
	var err error
	server.ID, err = database.Exec("INSERT INTO server (name, image) Values (?, ?);", args)
	if err != nil {
		fmt.Println("failed to add server")
	}
	return server
}

// CreateConnection .
func (s *Server) NewConnection(accountID int, alias string, permissions uint8, ws *websocket.Conn) *Connection {
	var args []interface{}
	args = append(args, s.ID, accountID, alias, permissions)
	_, err := database.Exec("INSERT INTO connection (server_id, account_id, alias, permissions) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	c := &Connection{Server: s, Alias: alias, Permissions: permissions, ws: ws}
	return c
}

// Delete .
func (s *Server) Delete() {
	var args []interface{}
	args = append(args, s.ID)
	_, err := database.Exec("DELETE FROM server WHERE id=?", args)
	if err != nil {
		panic(err)
	}
}
