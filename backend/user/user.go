package user

import (
	"bytes"
	"errors"
	"nebula/database"
	"nebula/random"

	"golang.org/x/crypto/argon2"
)

const (
	saltLength           = 32
	hashedPasswordLength = 32
)

// User holds the json values for an account that are sent to a client
type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Avatar         string `json:"avatar"`
	Email          string `json:"email"`
	hashedPassword []byte
	salt           []byte
}

// Get .
func Get(credentials Credentials) (*User, error) {

	var args []interface{}
	args = append(args, credentials.Email)
	rows, err := database.Query("SELECT id, email, username, avatar, password, salt FROM Account WHERE email=?;", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Email, &u.Username, &u.Avatar, &u.hashedPassword, &u.salt)
		defer u.clearPassword()
		if u.isPasswordCorrect(credentials.Password) {
			return &u, nil
		}
	}
	return nil, errors.New("Failed to get user Information")
}

func (u *User) clearPassword() {
	// u.hashedPassword = random.GetSecureBytes(32)
	// u.salt = random.GetSecureBytes(32)
	u.hashedPassword = nil
	u.salt = nil
}

// IsPasswordCorrect checks to see if the password in the database matches the password used to login
func (u *User) isPasswordCorrect(passwordAttempt string) bool {
	//storedPasswordHash, salt := getPassword(c.Email)
	hash := argon2.IDKey([]byte(passwordAttempt), u.salt, 4, 32*1024, 4, 32)
	return bytes.Compare(u.hashedPassword, hash) == 0
}

// Add adds an account to the database
func Add(a Account) *User {
	salt := random.GetSecureBytes(32)
	hashedPassword := argon2.IDKey([]byte(a.Password), salt, 4, 32*1024, 4, 32)
	var args []interface{}
	args = append(args, a.Email, a.Username, hashedPassword, salt, a.Avatar)
	rows, err := database.Query("INSERT INTO Account (email, username, password, salt, avatar) Values (?, ?, ?, ?, ?);", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	return &User{Username: a.Username, Email: a.Email, Avatar: a.Avatar}
}

// HasPermission .
func (u *User) HasPermission(permission, serverID int) bool {
	var args []interface{}
	args = append(args, u.ID)
	rows, err := database.Query(
		`SELECT permissions
			FROM Role
			INNER JOIN ServerMember ON Role.id = ServerMember.role_id
			where ServerMember.account_id=?;`, args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		var p int
		rows.Scan(&p)
		return p&permission == permission
	}
	return false
}
