package server

// Role .
type Role struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ServerPermissions uint8  `json:"serverPermissions"`
}
