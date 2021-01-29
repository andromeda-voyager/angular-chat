package server

import (
	"fmt"
	"nebula/database"
	"nebula/permissions"
)

// Channel .
type Channel struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}

// NewChannel .
func NewChannel(channel Channel, rolesWithAccess []Role, serverID int) *Channel {
	var args []interface{}
	args = append(args, serverID, channel.Name)
	fmt.Println(channel.Name)
	channelID, err := database.Exec("INSERT INTO Channel (server_id, name) Values (?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	var c = &Channel{ID: channelID, Name: channel.Name, Posts: nil}
	// ok := validateRoles(rolesWithAccess, serverID)
	// if ok {
	c.AddPermissions(rolesWithAccess)
	//	}
	return c
}

// getChannels .
func (c Channel) getPosts() {
	var args []interface{}
	args = append(args, c.ID)
	rows, err := database.Query(
		`SELECT text, media, time_posted, account_id
		FROM Post 
		where Post.channel_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p *Post
		rows.Scan(&p.Text, p.Media, p.TimePosted, p.AccountID)
		c.Posts = append(c.Posts, p)
	}
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
func (c *Channel) AddPermissions(roles []Role) {
	for _, r := range roles {
		var args []interface{}
		args = append(args, r.ID, c.ID, permissions.None)
		_, err := database.Exec("INSERT INTO ChannelPermissions (role_id, channel_id, permissions) Values (?, ?, ?);", args)
		if err != nil {
			panic(err.Error())
		}
	}
}
