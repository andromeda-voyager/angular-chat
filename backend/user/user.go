package user

import (
	"bytes"
	"errors"
	"nebula/database"
	"nebula/permissions"
	"nebula/util"

	"golang.org/x/crypto/argon2"
)

// User holds the json values for an account that are sent to a client
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	email    string
	password []byte
	salt     []byte
}

// IsPasswordCorrect checks to see if the password in the database matches the password used to login
func (u *User) IsPasswordCorrect(password string) bool {
	//storedPasswordHash, salt := getPassword(c.Email)
	hash := argon2.IDKey([]byte(password), []byte(u.salt), 4, 32*1024, 4, 32)
	return bytes.Compare(u.password, hash) == 0
}

// DeleteServer .
func (u *User) DeleteServer(serverID int) bool {
	var args []interface{}
	args = append(args, serverID, serverID)
	rows, err := database.Query(
		`SELECT server_permissions
			FROM Role
			INNER JOIN ServerMember ON Role.id = ServerMember.role_id
			where ServerMember.server_id=? AND Role.server_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p uint8
		rows.Scan(p)
		if permissions.CanDeleteServer(p) {

		}
	}
	return false
}

// Get .
func Get(email string) (*User, error) {
	var args []interface{}
	args = append(args, email)
	rows, err := database.Query("SELECT id, email, username, avatar, password, salt FROM Account WHERE email=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.email, &u.Username, &u.Avatar, &u.password, &u.salt)
		// a.Connections = getConnections(a.ID)
		return &u, nil
	}

	return nil, errors.New("Failed to get user Information")
}

// Add adds an account to the database
func Add(a Account) *User {
	salt := util.GetRandomBytes(32)
	hashedPassword := argon2.IDKey([]byte(a.Password), salt, 4, 32*1024, 4, 32)
	var args []interface{}
	args = append(args, a.Email, a.Username, hashedPassword, salt, a.Avatar)
	rows, err := database.Query("INSERT INTO Account (email, username, password, salt, avatar) Values (?, ?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	return &User{Username: a.Username, email: a.Email, Avatar: a.Avatar, password: hashedPassword, salt: salt}
}

// func (a *Account) CanPostOnServer(serverID int) bool {
// 	for _, v := range a.Connections {
// 		if v.Server.ID == serverID {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (a *Account) DeleteServer(serverID int) bool {
// 	for _, v := range a.Connections {
// 		if v.Server.ID == serverID {
// 			if permissions.CanDeleteServer(v.Permissions) {
// 				v.Server.Delete()
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // CreateConnection adds a connection to a user
// func (a *Account) AddConnection(c *server.Connection) {
// 	a.Connections = append(a.Connections, c)
// }
