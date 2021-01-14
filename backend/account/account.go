package account

import (
	"database/sql"
	"nebula/config"
	"nebula/permissions"
	"nebula/server"

	"golang.org/x/net/websocket"
)

// Account holds all the values stored for each user in the database. Used to unmarshal json sent during account creation.
type Account struct {
	Username    string               `json:"username"`
	Password    string               `json:"password"`
	Email       string               `json:"email"`
	AvatarURL   string               `json:"avatarURL"`
	Code        string               `json:"code"`
	Connections []*server.Connection `json:"connections"`
	ID          int                  `json:"id"`
	ws          *websocket.Conn
}

// GetConnection .
func (a *Account) GetConnection(index int) *server.Connection {
	if len(a.Connections) > index {
		return a.Connections[index]
	}
	return nil
}

func (a *Account) hasValidFields() bool {
	if len(a.Email) < 5 {
		return false
	}
	if len(a.Username) < 1 {
		return false
	}
	if len(a.Password) < 8 {
		return false
	}
	return true
}

func (a *Account) CanPostToServer(serverID int) bool {
	for _, v := range a.Connections {
		if v.Server.ID == serverID {
			return true
		}
	}
	return false
}

func (a *Account) DeleteServer(serverID int) bool {
	for _, v := range a.Connections {
		if v.Server.ID == serverID {
			if permissions.CanDeleteServer(v.Permissions) {
				v.Server.Delete()
				return true
			}
		}
	}
	return false
}

// AddConnection adds a connection to a user
func (a *Account) CreateConnection(s server.Server, permissions uint8) {
	c := s.NewConnection(a.ID, a.Username, permissions, a.ws)
	a.Connections = append(a.Connections, c)
}

// IsEmailInUse checks if an email is already used for an account
func IsEmailInUse(email string) bool {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM account WHERE email=?;", email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	return rows.Next()
}

// GetServerID .
// func (a *Account) GetServerID(index int) int {
// 	if len(a.Connections) > index {
// 		return a.Connections[index].ServerID
// 	}
// 	return -1
// }
