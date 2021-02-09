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

func Get(serverID, userID int) (*Server, bool) {
	var args []interface{}
	args = append(args, userID, serverID)
	rows, err := database.Query(
		`SELECT Server.id, Server.name, Server.image, Server.description, 
		ServerMember.alias,
		Role.id, Role.name, Role.permissions
		FROM Server 
		INNER JOIN ServerMember ON Server.id = ServerMember.server_id 
		INNER JOIN Role ON ServerMember.role_id = Role.id 
		where ServerMember.account_id=? AND Server.id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		var s Server
		var r Role
		rows.Scan(&s.ID, &s.Name, &s.Image, &s.Description, &s.Alias, &r.ID, &r.Name, &r.Permissions)
		s.Role = r
		s.LoadRoles()
		s.LoadChannels()
		return &s, true
	}
	return nil, false
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
	roleID, err := database.Exec("INSERT INTO Role (server_id, name, ranking, permissions) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	r := Role{ID: roleID, Name: name}
	s.Roles = append(s.Roles, r)
	return r
}

// Delete .
func Delete(id int) bool {
	var args []interface{}
	args = append(args, id)
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
	s.Channels = []Channel{}
	for rows.Next() {
		var c Channel
		rows.Scan(&c.ID, &c.Name)
		c.LoadOverrides()
		s.Channels = append(s.Channels, c)
	}
}

// LoadRoles .
func (s *Server) LoadRoles() {
	var args []interface{}
	args = append(args, s.ID)
	rows, err := database.Query(
		`SELECT id, name, ranking, permissions
		FROM Role
		where server_id=?
		ORDER BY
		ranking ASC;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	s.Roles = []Role{}
	for rows.Next() {
		var r Role
		rows.Scan(&r.ID, &r.Name, &r.Ranking, &r.Permissions)
		s.Roles = append(s.Roles, r)
	}
}

func getServers(accountID int) []*Server {
	var servers []*Server = []*Server{}
	var args []interface{}
	args = append(args, accountID)
	rows, err := database.Query(
		`SELECT Server.id, Server.name, Server.image, Server.description, 
		ServerMember.alias,
		Role.id, Role.name, Role.permissions
		FROM Server 
		INNER JOIN ServerMember ON Server.id = ServerMember.server_id 
		INNER JOIN Role ON ServerMember.role_id = Role.id 
		where ServerMember.account_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var s Server
		var r Role
		rows.Scan(&s.ID, &s.Name, &s.Image, &s.Description, &s.Alias, &r.ID, &r.Name, &r.Permissions)
		s.Role = r
		s.LoadRoles()
		s.LoadChannels()
		servers = append(servers, &s)
	}

	return servers

}
