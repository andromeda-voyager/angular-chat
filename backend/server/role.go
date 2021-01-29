package server

import "nebula/database"

// Role .
type Role struct {
	ID                 int                  `json:"id"`
	Ranking            int                  `json:"ranking"`
	Name               string               `json:"name"`
	ServerPermissions  uint8                `json:"serverPermissions"`
	ChannelPermissions []ChannelPermissions `json:"channelPermissions"`
}

// ChannelPermissions .
type ChannelPermissions struct {
	ChannelID int   `json:"channelID"`
	Value     uint8 `json:"value"`
}

// LoadChannelPermissions .
func (r Role) LoadChannelPermissions() {
	var args []interface{}
	args = append(args, r.ID)
	rows, err := database.Query(
		`SELECT channel_id, permissions
		FROM ChannelPermissions
		where role_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var permissions []ChannelPermissions
	for rows.Next() {
		var p ChannelPermissions
		rows.Scan(&p.ChannelID, &p.Value)
		permissions = append(permissions, p)
	}
}
