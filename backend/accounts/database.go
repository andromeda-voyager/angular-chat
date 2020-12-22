package accounts

import (
	"database/sql"
	"nebula/config"
	"nebula/util"

	"golang.org/x/crypto/argon2"
)

// Add user to the database
func Add(user User) {
	salt := util.GetRandomBytes(32)
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 4, 32*1024, 4, 32)
	//	password64 := base64.RawStdEncoding.EncodeToString(hashedPassword)
	db, err := sql.Open("mysql", config.Username+":"+config.Password+"@tcp(localhost:3306)/nebula")
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

// DoesAccountExist checks if email is already used for an account
func DoesAccountExist(email string) bool {
	return false
}

func getPassword(email string) ([]byte, []byte) {
	db, err := sql.Open("mysql", config.Username+":"+config.Password+"@tcp(localhost:3306)/nebula")
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
