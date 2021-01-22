package server

import (
	"time"
)

// Post .
type Post struct {
	Text       string    `json:"text"`
	Media      string    `json:"Media"`
	TimePosted time.Time `json:"timePosted"`
	AccountID  int       `json:"accountID"`
	// User       user.User `json:"user"`
}
