package main

// Connection .
type Connection struct {
	ServerID    int    `json:"serverID"`
	UserID      int    `json:"userID"`
	Alias       string `json:"alias"`
	Permissions uint8  `json:"permissions"`
}
