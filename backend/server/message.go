package server

import (
	"nebula/database"
	"time"
)

type MessageUpdate struct {
	Type    string  `json:"type"`
	Event   string  `json:"event"`
	Message Message `json:"message"`
}

// Message .
type Message struct {
	ID         int       `json:"id"`
	ChannelID  int       `json:"channelID"`
	Text       string    `json:"text"`
	Media      string    `json:"Media"`
	TimePosted time.Time `json:"timePosted"`
	Member     Member    `json:"member"`
}

func (m *Message) Add(senderID int) (bool, error) {
	var args []interface{}
	m.TimePosted = time.Now().UTC()
	m.Member = GetMember(senderID)
	args = append(args, m.Member.AccountID, m.ChannelID, m.Text, m.Media, m.TimePosted)
	_, err := database.Exec("INSERT INTO Message (account_id, channel_id, text, media, time_posted) Values (?, ?, ?, ?, ?);", args)
	if err != nil {
		return false, err
	}
	return true, err
}

func deleteMessage(messageID int) (bool, error) {
	var args []interface{}
	args = append(args, messageID)
	_, err := database.Exec("DELETE FROM Message WHERE id=?", args)
	if err != nil {
		return false, err
	}
	return true, nil
}
