package server

// Server .
type Server struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
	Posts    []Post `json:"posts"`
}

// Invite .
type Invite struct {
	Code     string `json:"code"`
	ServerID string `json:"serverID"`
}
