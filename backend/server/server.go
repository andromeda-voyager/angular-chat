package server

import (
	"encoding/json"
	"fmt"
	"nebula/database"
	"nebula/permissions"
	"nebula/util"
	"net/http"
)

// Server .
type Server struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Role        *Role      `json:"role"`
	Roles       []*Role    `json:"roles"`
	Alias       string     `json:"alias"`
	Channels    []*Channel `json:"channels"`
}

// Member .
type Member struct {
	AccountID int    `json:"accountID"`
	Alias     string `json:"alias"`
	Avatar    string `json:"avatar"`
}

// Invite .
type Invite struct {
	Code     string `json:"code"`
	ServerID string `json:"serverID"`
}

// New .
func New(m *Member, r *http.Request) Server {
	serverJSON := r.FormValue("server")
	var server Server
	json.Unmarshal([]byte(serverJSON), &server)
	server.Image = util.SaveImage(r)
	var args []interface{}
	args = append(args, server.Name, server.Image, server.Description)
	fmt.Println(server.Name)
	var err error
	server.ID, err = database.Exec("INSERT INTO Server (name, image, description) Values (?, ?, ?);", args)
	if err != nil {
		fmt.Println("failed to add server")
	}
	server.NewRole("owner", permissions.Full)
	server.NewRole("default", permissions.None)

	server.NewMember(m)
	return server
}

// NewMember .
func (s *Server) NewMember(m *Member) {
	var args []interface{}
	args = append(args, s.ID, m.AccountID, m.Alias, s.Role.ID)
	_, err := database.Exec("INSERT INTO ServerMember (server_id, account_id, alias, role_id) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	s.Alias = m.Alias
}

// NewRole .
func (s *Server) NewRole(name string, permissions uint8) {
	var args []interface{}
	args = append(args, s.ID, name, permissions)
	roleID, err := database.Exec("INSERT INTO Role (server_id, name, server_permissions) Values (?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	role := &Role{ID: roleID, Name: name}
	s.Role = role
}

// Delete .
func (s *Server) Delete() bool {
	var args []interface{}
	args = append(args, s.ID)
	_, err := database.Exec("DELETE FROM Server WHERE id=?", args)
	if err != nil {
		return false
	}
	return true
}

// GetChannels .
func (s *Server) GetChannels() {
	fmt.Println("Getting channels")
	var args []interface{}
	args = append(args, s.Role.ID)
	rows, err := database.Query(
		`SELECT ChannelPermissions.permissions, Channel.id, Channel.name
		FROM ChannelPermissions 
		INNER JOIN Channel ON ChannelPermissions.channel_id = Channel.id 
		where ChannelPermissions.role_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var c *Channel
		rows.Scan(&c.ID, &c.Name)
		s.Channels = append(s.Channels, c)
	}
}
