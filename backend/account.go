package main

import (
	"nebula/database"
)

// Account holds all the values stored for each user in the database. Used to unmarshal json sent during account creation.
type Account struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatarURL"`
	Code        string `json:"code"`
	Connections []int  `json:"servers"`
	id          int
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
	a.Connections = append(a.Connections, serverID)
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
