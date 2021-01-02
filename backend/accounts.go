package main

import (
	"errors"
	"nebula/database"
	"nebula/util"

	"golang.org/x/crypto/argon2"
)

func getUser(email string) (*Account, error) {
	var args []interface{}
	args = append(args, email)
	rows, err := database.Query("SELECT Email, Username, AvatarURL FROM Users WHERE Email=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var a Account
		rows.Scan(&a.Email, &a.Username, &a.AvatarURL)
		return &a, nil
	}
	return nil, errors.New("Failed to get user Information")
}

// addAccount user to the database
func addAccount(a Account) {
	salt := util.GetRandomBytes(32)
	hashedPassword := argon2.IDKey([]byte(a.Password), salt, 4, 32*1024, 4, 32)
	var args []interface{}
	args = append(args, a.Email, a.Username, hashedPassword, salt, a.AvatarURL)
	rows, err := database.Query("INSERT INTO Users (Email, Username, Password, Salt, AvatarURL) Values (?, ?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
}

func getPassword(email string) ([]byte, []byte) {

	var args []interface{}
	args = append(args, email)
	rows, err := database.Query("SELECT Password, Salt FROM Users WHERE Email=?;", args)
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
