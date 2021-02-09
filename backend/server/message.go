package server

import (
	"nebula/database"
	"time"
)

// Post .
type Message struct {
	ID         int       `json:"id"`
	AccountID  int       `json:"AccountID"`
	ChannelID  string    `json:"channelID"`
	Text       string    `json:"text"`
	Media      string    `json:"Media"`
	TimePosted time.Time `json:"timeSent"`
	Member     Member    `json:"member"`
}

func NewMessage(m *Message) {
	var args []interface{}
	args = append(args, m.AccountID, m.ChannelID, m.Text, m.Media)
	id, err := database.Exec("INSERT INTO Message (account_id, channel_id, text, media) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	m.ID = id
}
