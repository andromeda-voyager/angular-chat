package server

import (
	"nebula/database"
)

// Member .
type Member struct {
	AccountID int    `json:"accountID"`
	Alias     string `json:"alias"`
	Avatar    string `json:"avatar"`
	Role      Role   `json:"role"`
}

func getMember(accountID int) Member {

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
