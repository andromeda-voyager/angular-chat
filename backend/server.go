package main

import (
	"nebula/database"
)

// Server .
type Server struct {
	ID             int
	Name           string `json:"name"`
	ServerImageURL string `json:"serverImageURL"`
}

// AddServer user to the database
func addServer(server Server) (int, error) {
	var args []interface{}
	args = append(args, server.Name, server.ServerImageURL)
	return database.Exec("INSERT INTO Servers (Name, ServerImageURL) Values (?, ?);", args)
}
