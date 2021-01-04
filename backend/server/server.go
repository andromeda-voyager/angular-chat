package server

// Server .
type Server struct {
	ID             int
	Name           string `json:"name"`
	ServerImageURL string `json:"serverImageURL"`
}

// Invite .
type Invite struct {
	Code     string `json:"code"`
	ServerID string `json:"serverID"`
}
