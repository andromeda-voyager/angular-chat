package account

import (
	"database/sql"
	"nebula/config"
	"nebula/database"
	"nebula/server"
)

// Account holds all the values stored for each user in the database. Used to unmarshal json sent during account creation.
type Account struct {
	Username    string               `json:"username"`
	Password    string               `json:"password"`
	Email       string               `json:"email"`
	AvatarURL   string               `json:"avatarURL"`
	Code        string               `json:"code"`
	Connections []*server.Connection `json:"connections"`
	id          int
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

// AddConnection adds a connection to a user
func (a *Account) AddConnection(serverID int, permissions uint8) {
	var args []interface{}
	args = append(args, serverID, a.id, a.Username, permissions)
	_, err := database.Exec("INSERT INTO Connections (ServerID, UserID, Alias, Permissions) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
}

// ID returns the id of the account
func (a *Account) ID() int {
	return a.id
}

// IsEmailInUse checks if an email is already used for an account
func IsEmailInUse(email string) bool {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM Users WHERE Email=?;", email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	return rows.Next()
}

// GetServerID .
func (a *Account) GetServerID(index int) int {
	if len(a.Connections) > index {
		return a.Connections[index].ServerID
	}
	return -1
}
