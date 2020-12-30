package accounts

import (
	"database/sql"
	"errors"
	"fmt"
	"nebula/config"
	"nebula/util"

	"golang.org/x/crypto/argon2"
)

func isValidUser(user User) (bool, string) {
	if len(user.Name) < 1 {
		return false, "Invalid Name"
	}
	if len(user.Email) < 5 {
		return false, "Invalid Email"
	}
	if len(user.Username) < 1 {
		return false, "Invalid Username"
	}
	if len(user.Password) < 8 {
		return false, "Invalid Password"
	}
	return true, ""
}

func TestQuery() {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT Email FROM Users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	if rows.Next() {
		var user User
		rows.Scan(&user.Email)
		fmt.Println(user.Email)
	}
}

func Get(email string) (*User, error) {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT Email, Name, AvatarURL FROM Users WHERE Email=?;", email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	if rows.Next() {
		var user User
		rows.Scan(&user.Email, &user.Name, &user.AvatarURL)
		return &user, nil
	}
	return nil, errors.New("Failed to get user Information")
}

// Add user to the database
func Add(user User) {
	salt := util.GetRandomBytes(32)
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 4, 32*1024, 4, 32)
	//	password64 := base64.RawStdEncoding.EncodeToString(hashedPassword)
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("INSERT INTO Users (Email, Name, Password, Salt, AvatarURL) Values (?, ?, ?, ?, ?);", user.Email, user.Name, hashedPassword, salt, user.AvatarURL)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()
}

func getPassword(email string) ([]byte, []byte) {
	db, err := sql.Open("mysql", config.DatabaseUser+":"+config.DatabasePassword+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT Password, Salt FROM Users WHERE Email=?;", email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	if rows.Next() {
		var password []byte
		var salt = make([]byte, 32)
		rows.Scan(&password, &salt)
		return password, salt
	}
	return nil, nil
}
