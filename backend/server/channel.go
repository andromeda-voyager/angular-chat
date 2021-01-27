package server

import "nebula/database"

// Channel .
type Channel struct {
	ID                 int                  `json:"id"`
	Name               string               `json:"name"`
	Posts              []*Post              `json:"posts"`
	ChannelPermissions []ChannelPermissions `json:"channelPermissions"`
}

// NewChannel .
type NewChannel struct {
	ServerID           int                  `json:"serverID"`
	Name               string               `json:"name"`
	ChannelPermissions []ChannelPermissions `json:"channelPermissions"`
}

// ChannelPermissions .
type ChannelPermissions struct {
	RoleID      int
	RoleRank    int   `json:"roleRank"`
	Permissions uint8 `json:"permissions"`
}

// getChannels .
func (c *Channel) getPosts() {
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

// AddChannelPermissions .
func (c *Channel) AddChannelPermissions(channelPermissions ChannelPermissions) {

	var args []interface{}
	args = append(args, channelPermissions.RoleID, c.ID, channelPermissions.Permissions)
	_, err := database.Exec("INSERT INTO ChannelPermissions (role_id, channel_id, permissions) Values (?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}

}
