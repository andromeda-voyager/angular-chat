package server

import (
	"encoding/json"
	"fmt"
	"nebula/database"
	"nebula/images"
	"nebula/permissions"
	"nebula/user"
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
	Members     []Member  `json:"members"`
	Channels    []Channel `json:"channels"`
	Invites     []Invite  `json:"invites"`
}

const DefaultImageUrl = "default-avatar.jpg"

// New .
func New(owner user.User, r *http.Request) Server {
	serverJSON := r.FormValue("server")
	var s Server
	json.Unmarshal([]byte(serverJSON), &s)
	s.Image = images.Save(r, DefaultImageUrl)
	var args []interface{}
	args = append(args, s.Name, s.Image, s.Description)
	var err error
	s.ID, err = database.Exec("INSERT INTO Server (name, image, description) Values (?, ?, ?);", args)
	if err != nil {
		fmt.Println("failed to add server to database")
	}
	ownerRole := NewRole(s.ID, "owner", 0, permissions.Full)
	s.Roles = append(s.Roles, ownerRole)
	s.Roles = append(s.Roles, NewRole(s.ID, "default", 1, permissions.None)) // default role is created for new members
	m := NewMember(s.ID, owner, ownerRole)
	//s.Members = []Member{}
	s.Alias = owner.Username
	s.Members = append(s.Members, m)
	return s
}

func Get(serverID int, userID int) (*Server, bool) {
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

		s.LoadInvites()
		return &s, true
	}
	return nil, false
}

// Join .
func Join(u user.User, invite Invite) (*Server, bool) {
	var args []interface{}
	args = append(args, invite.Code, "default")
	rows, err := database.Query(`SELECT Invite.server_id, Role.id, Role.ranking, Role.name, Role.permissions
		FROM Invite
		INNER JOIN Role ON Invite.server_id = Role.server_id
		WHERE Invite.code=? and Role.name=?;`, args)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var serverID int
		var r Role
		rows.Scan(&serverID, &r.ID, &r.Ranking, &r.Name, &r.Permissions)

		m := NewMember(serverID, u, r)
		s, ok := Get(serverID, m.AccountID)
		if ok {
			return s, true
		}

	}
	return nil, false
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

// LoadInvites .
func (s *Server) LoadInvites() {
	var args []interface{}
	args = append(args, s.ID)
	rows, err := database.Query(
		`SELECT code
		FROM Invite
		where server_id=?`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	s.Invites = []Invite{}
	for rows.Next() {
		var i Invite
		rows.Scan(&i.Code)
		s.Invites = append(s.Invites, i)
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
		s.LoadInvites()
		servers = append(servers, &s)
	}

	return servers

}
