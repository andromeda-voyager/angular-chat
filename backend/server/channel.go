package server

import "nebula/database"

// Channel .
type Channel struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Permissions uint8   `json:"permissions"`
	Posts       []*Post `json:"posts"`
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
