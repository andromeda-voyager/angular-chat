package server

import (
	"nebula/database"
	"nebula/permissions"
)

// Channel .
type Channel struct {
	ID        int        `json:"id"`
	ServerID  int        `json:"serverID"`
	Name      string     `json:"name"`
	Overrides []Override `json:"overrides"`
}

// Override .
type Override struct {
	RoleID      int   `json:"roleID"`
	Permissions uint8 `json:"permissions"`
}

// NewChannel .
func NewChannel(c *Channel) {
	var args []interface{}
	args = append(args, c.ServerID, c.Name)
	channelID, err := database.Exec("INSERT INTO Channel (server_id, name) Values (?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	c.ID = channelID
	c.AddOverrides()

}

// LoadChannelOverrides .
func (c Channel) LoadOverrides() {
	var args []interface{}
	args = append(args, c.ID)
	rows, err := database.Query(
		`SELECT role_id, permissions
		FROM ChannelPermissions
		where channel_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	c.Overrides = []Override{}
	for rows.Next() {
		var o Override
		rows.Scan(&o.RoleID, &o.Permissions)
		c.Overrides = append(c.Overrides, o)
	}
}

// getChannels .
func (c Channel) getPosts() {
	// var args []interface{}
	// args = append(args, c.ID)
	// rows, err := database.Query(
	// 	`SELECT text, media, time_sent, account_id
	// 	FROM Message
	// 	where Message.channel_id=?;`, args)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var m Message
	// 	rows.Scan(&m.Text, &m.Media, &m.TimePosted, &m.AccountID)
	// 	c.Messages = append(c.Messages, &m)
	// }
}

// func validateRoles(rolesWithAccess []Role, serverID int) bool {
// 	allRoles := getServerRoles(serverID)
// 	for _, r := range rolesWithAccess {
// 		if r.Rank <= len(allRoles) {
// 			if r.ID != allRoles[r.Rank].ID {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

// AddPermissions .
func (c *Channel) AddOverrides() {
	for _, r := range c.Overrides {
		var args []interface{}
		args = append(args, r.RoleID, c.ID, permissions.None)
		_, err := database.Exec("INSERT INTO ChannelPermissions (role_id, channel_id, permissions) Values (?, ?, ?);", args)
		if err != nil {
			panic(err.Error())
		}
	}
}
