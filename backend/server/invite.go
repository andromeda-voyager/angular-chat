package server

import (
	"fmt"
	"nebula/database"
	"nebula/random"
)

// Invite .
type Invite struct {
	Code     string `json:"code"`
	ServerID int    `json:"serverID"`
}

// NewInvite .
func NewInvite(serverID int) Invite {
	var args []interface{}
	code := random.NewString(7)
	fmt.Println(code)
	args = append(args, serverID, code)
	_, err := database.Exec("INSERT INTO Invite (server_id, code) Values (?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	i := Invite{ServerID: serverID, Code: code}
	return i
}
