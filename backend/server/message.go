package server

import (
	"time"
)

// Post .
type Message struct {
	Text       string    `json:"text"`
	Media      string    `json:"Media"`
	TimePosted time.Time `json:"timeSent"`
	AccountID  int       `json:"accountID"`
	Member     Member    `json:"member"`
}
