package servers

import (
	"database/sql"
	"nebula/config"
)

// Add user to the database
func Add(server Server) {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("INSERT INTO Servers (Name, ServerImageURL) Values (?, ?);", server.Name, server.ServerImageURL)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()
}
