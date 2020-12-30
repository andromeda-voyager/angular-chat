package servers

// Server .
type Server struct {
	ServerID       string
	Name           string `json:"name"`
	ServerImageURL string `json:"serverImageURL"`
}
