package accounts

import (
	"database/sql"
	"nebulous/config"
	"nebulous/util"

	"golang.org/x/crypto/argon2"
)

func addUser(user User) {
	salt := util.GetRandomBytes(32)
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 4, 32*1024, 4, 32)
	//	password64 := base64.RawStdEncoding.EncodeToString(hashedPassword)
	db, err := sql.Open("mysql", config.Username+":"+config.Password+"@tcp(localhost:3306)/nebula")
	if err != nil {
		panic(err.Error())
	}
	//"localhost:8080/avatars/mp.jpg"
	rows, err := db.Query("INSERT INTO Users (Email, Name, Password, Salt, AvatarURL) Values (?, ?, ?, ?, ?);", user.Email, user.Name, hashedPassword, salt, user.AvatarURL)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()
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
	// for rows.Next() {
	// 	var name string
	// 	if err := rows.Scan(&name); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(name)
	// }

}
