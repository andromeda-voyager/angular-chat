package main

import (
	"errors"
	"nebula/account"
	"nebula/database"
	"nebula/server"
	"nebula/util"

	"golang.org/x/crypto/argon2"
)

func getAccount(email string) (*account.Account, error) {
	var args []interface{}
	args = append(args, email)
	rows, err := database.Query("SELECT AccountID, Email, Username, AvatarURL FROM Accounts WHERE Email=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var a account.Account
		rows.Scan(&a.ID, &a.Email, &a.Username, &a.AvatarURL)
		a.Connections = getConnections(a.ID)
		return &a, nil
	}

	return nil, errors.New("Failed to get user Information")
}

func getServer(serverID int) *server.Server {
	var args []interface{}
	args = append(args, serverID)
	rows, err := database.Query("SELECT ServerID, Name, ImageURL FROM Servers WHERE ServerID=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var s server.Server
	for rows.Next() {
		rows.Scan(&s.ID, &s.Name, &s.ImageURL)
		s.Posts = getPosts(s.ID)
	}
	return &s
}

func getPosts(serverID int) []server.Post {
	var args []interface{}
	args = append(args, serverID)
	rows, err := database.Query("SELECT Text, MediaURL, TimePosted, AccountID FROM Posts WHERE ServerID=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var posts []server.Post
	for rows.Next() {
		var p server.Post
		var id int
		rows.Scan(&p.Text, &p.MediaURL, &p.TimePosted, &id)
		posts = append(posts, p)
	}
	return posts
}

func getConnections(accountID int) []server.Connection {
	var args []interface{}
	args = append(args, accountID)
	rows, err := database.Query("SELECT ServerID, Alias, Permissions FROM Connections WHERE AccountID=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var connections []server.Connection
	for rows.Next() {
		var c server.Connection
		var serverID int
		rows.Scan(&serverID, &c.Alias, &c.Permissions)
		c.Server = getServer(serverID)
		connections = append(connections, c)
	}
	return connections
}

// addAccount user to the database
func addAccount(a account.Account) {
	salt := util.GetRandomBytes(32)
	hashedPassword := argon2.IDKey([]byte(a.Password), salt, 4, 32*1024, 4, 32)
	var args []interface{}
	args = append(args, a.Email, a.Username, hashedPassword, salt, a.AvatarURL)
	rows, err := database.Query("INSERT INTO Accounts (Email, Username, Password, Salt, AvatarURL) Values (?, ?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
}

func getPassword(email string) ([]byte, []byte) {

	var args []interface{}
	args = append(args, email)
	rows, err := database.Query("SELECT Password, Salt FROM Accounts WHERE Email=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var password []byte
		var salt = make([]byte, 32)
		rows.Scan(&password, &salt)
		return password, salt
	}
	return nil, nil
}
