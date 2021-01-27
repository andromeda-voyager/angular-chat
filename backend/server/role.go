package server

import "nebula/database"

// Role .
type Role struct {
	ID                int    `json:"id"`
	Rank              int    `json:"rank"`
	Name              string `json:"name"`
	ServerPermissions uint8  `json:"serverPermissions"`
}

// getRoles .
func getRoles(serverID int) []Role {
	var args []interface{}
	args = append(args, serverID)
	rows, err := database.Query(
		`SELECT id, rank
		FROM Role
		ORDER BY
		rank ASC
		where server_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var roles []Role
	for rows.Next() {
		var r Role
		rows.Scan(&r.ID, &r.Rank)
		roles = append(roles, r)
	}
	return roles
}
