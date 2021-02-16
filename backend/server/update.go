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
	Type    string  `json:"type"`
	Event   string  `json:"event"`
	Server  Server  `json:"server,omitempty"`
	Channel Channel `json:"channel,omitempty"`
	Role    Role    `json:"role,omitempty"`
	Message Message `json:"message,omitempty"`
}

type MessageUpdate struct {
	Type    string  `json:"type"`
	Event   string  `json:"event"`
	Message Message `json:"message"`
}
