package server

import (
	"time"
)

// Post .
type Post struct {
	Text       string    `json:"text"`
	MediaURL   string    `json:"mediaURL"`
	ServerID   int       `json:"serverID"`
	TimePosted time.Time `json:"timePosted"`
	UserID     int       `json:"userID"`
}
