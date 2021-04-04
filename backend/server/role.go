package server

import "nebula/database"

// Role .
type Role struct {
	ID          int    `json:"id"`
	Ranking     int    `json:"ranking"`
	Name        string `json:"name"`
	Permissions uint8  `json:"permissions"`
}

func GetRole(serverID int, name string) (Role, bool) {
	var args []interface{}
	args = append(args, serverID, name)
	rows, err := database.Query(
		`SELECT id, ranking, permissions
		FROM Role 
		where server_id=? AND name=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var r Role
	if rows.Next() {
		var r Role
		r.Name = name
		rows.Scan(&r.ID, &r.Ranking, &r.Permissions)
		return r, true
	}
	return r, false
}

// NewRole .
func NewRole(serverID int, name string, ranking int, permissions uint8) Role {
	var args []interface{}
	args = append(args, serverID, name, ranking, permissions)
	roleID, err := database.Exec("INSERT INTO Role (server_id, name, ranking, permissions) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	r := Role{ID: roleID, Name: name}
	return r
}
