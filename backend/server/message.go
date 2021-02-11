package server

import (
	"nebula/database"
	"time"
)

// Post .
type Message struct {
	ID         int       `json:"id"`
	AccountID  int       `json:"AccountID"`
	ChannelID  int       `json:"channelID"`
	Text       string    `json:"text"`
	Media      string    `json:"Media"`
	TimePosted time.Time `json:"timePosted"`
	Member     Member    `json:"member"`
}

func (m *Message) Add() {
	var args []interface{}
	m.TimePosted = time.Now().UTC()
	args = append(args, m.AccountID, m.ChannelID, m.Text, m.Media, m.TimePosted)
	id, err := database.Exec("INSERT INTO Message (account_id, channel_id, text, media, time_posted) Values (?, ?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	m.ID = id
}
