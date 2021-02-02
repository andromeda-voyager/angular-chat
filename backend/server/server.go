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
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Role        Role      `json:"role"`
	Roles       []Role    `json:"roles"`
	Alias       string    `json:"alias"`
	Channels    []Channel `json:"channels"`
}

// Member .
type Member struct {
	AccountID int    `json:"accountID"`
	Alias     string `json:"alias"`
	Avatar    string `json:"avatar"`
	Role      Role   `json:"role"`
}

// Invite .
type Invite struct {
	Code     string `json:"code"`
	ServerID string `json:"serverID"`
}

// New .
func New(m *Member, r *http.Request) Server {
	serverJSON := r.FormValue("server")
	var s Server
	json.Unmarshal([]byte(serverJSON), &s)
	s.Image = util.SaveImage(r)
	var args []interface{}
	args = append(args, s.Name, s.Image, s.Description)
	var err error
	s.ID, err = database.Exec("INSERT INTO Server (name, image, description) Values (?, ?, ?);", args)
	if err != nil {
		fmt.Println("failed to add server")
	}
	s.Role = s.NewRole("owner", 0, permissions.Full)
	s.NewRole("default", 1, permissions.None)

	s.NewMember(m)
	return s
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
func (s *Server) NewRole(name string, ranking int, permissions uint8) Role {
	var args []interface{}
	args = append(args, s.ID, name, ranking, permissions)
	roleID, err := database.Exec("INSERT INTO Role (server_id, name, ranking, server_permissions) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	r := Role{ID: roleID, Name: name}
	s.Roles = append(s.Roles, r)
	return r
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

// LoadChannels .
func (s *Server) LoadChannels() {
	var args []interface{}
	args = append(args, s.Role.ID)
	rows, err := database.Query(
		`SELECT Channel.id, Channel.name
		FROM ChannelPermissions 
		INNER JOIN Channel ON ChannelPermissions.channel_id = Channel.id 
		where ChannelPermissions.role_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var c Channel
		rows.Scan(&c.ID, &c.Name)
		s.Channels = append(s.Channels, c)
	}
}

// LoadRoles .
func (s *Server) LoadRoles() {
	var args []interface{}
	args = append(args, s.ID)
	rows, err := database.Query(
		`SELECT id, name, ranking, server_permissions
		FROM Role
		where server_id=?
		ORDER BY
		ranking ASC;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var r Role
		rows.Scan(&r.ID, &r.Name, &r.Ranking, &r.ServerPermissions)
		r.LoadChannelPermissions()
		s.Roles = append(s.Roles, r)
	}
}
