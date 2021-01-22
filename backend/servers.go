package main

import (
	"fmt"
	"nebula/database"
	"nebula/server"
)

func getServers(accountID int) []*server.Server {
	var servers []*server.Server
	var args []interface{}
	fmt.Println("id: ", accountID)
	args = append(args, accountID)
	rows, err := database.Query(
		`SELECT Server.id, Server.name, Server.image, Server.description, 
		ServerMember.alias,
		Role.id, Role.name, Role.server_permissions
		FROM Server 
		INNER JOIN ServerMember ON Server.id = ServerMember.server_id 
		INNER JOIN Role ON ServerMember.role_id = Role.id 
		where ServerMember.account_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var s server.Server
		var r server.Role
		rows.Scan(&s.ID, &s.Name, &s.Image, &s.Description, &s.Alias, &r.ID, &r.Name, &r.ServerPermissions)
		s.Role = &r
		s.GetChannels()
		servers = append(servers, &s)
	}

	return servers

}
