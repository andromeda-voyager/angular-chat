package server

import (
	"nebula/database"
	"nebula/user"
)

// Member .
type Member struct {
	AccountID int    `json:"accountID"`
	Alias     string `json:"alias"`
	Avatar    string `json:"avatar"`
	Role      Role   `json:"role"`
}

func GetMember(accountID int) Member {

	var args []interface{}
	args = append(args, accountID)
	rows, err := database.Query(
		`SELECT alias, Account.avatar, Role.id, Role.ranking, Role.name, Role.permissions 
		FROM ServerMember
		INNER JOIN Role ON ServerMember.server_id = Role.server_id 
		INNER JOIN Account ON ServerMember.account_id = Account.id
		WHERE ServerMember.account_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var m Member = Member{AccountID: accountID}
	if rows.Next() {
		var r Role
		rows.Scan(&m.Alias, &m.Avatar, &r.ID, &r.Ranking, &r.Name, &r.Permissions)
	}
	return m
}

// NewMember .
func NewMember(serverID int, u user.User, r Role) Member {
	var args []interface{}
	args = append(args, serverID, u.ID, u.Username, r.ID)
	_, err := database.Exec("INSERT INTO ServerMember (server_id, account_id, alias, role_id) Values (?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	return Member{AccountID: u.ID, Alias: u.Username, Role: r, Avatar: u.Avatar}
}
