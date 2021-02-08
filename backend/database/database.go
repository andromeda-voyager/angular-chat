package database

import (
	"database/sql"
	"nebula/config"
)

func Query(stmt string, args []interface{}) (*sql.Rows, error) {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db.Query(stmt, args...)
}

func Exec(stmt string, args []interface{}) (int, error) {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		return -1, err
	}
	result, err := db.Exec(stmt, args...)
	if err != nil {
		return -1, err
	}
	serverID, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	return int(serverID), nil
}

func exec(stmt string) {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(stmt)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func GetServerFromInvite(code string) {

}
