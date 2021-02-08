package server

// Role .
type Role struct {
	ID          int    `json:"id"`
	Ranking     int    `json:"ranking"`
	Name        string `json:"name"`
	Permissions uint8  `json:"permissions"`
}
