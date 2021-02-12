package server

const (
	MESSAGE = "Message"
	CHANNEL = "Channel"
	ROLE    = "Role"
	MEMBER  = "Member"
	DELETE  = "Delete"
	MODIFY  = "Modify"
	NEW     = "New"
)

type Update struct {
	UpdateType string  `json:"updateType"`
	Event      string  `json:"event"`
	Server     Server  `json:"server"`
	Channel    Channel `json:"channel"`
	Role       Role    `json:"role"`
	Message    Message `json:"message"`
}
