package user

import (
	"fmt"
	"nebula/random"
	"nebula/router"
	"net/http"
	"time"
)

var loggedInUsers map[string]*User
var codes map[string]string

func init() {
	loggedInUsers = make(map[string]*User)
	codes = make(map[string]string)

}

func AddSession(user *User) *http.Cookie {
	token := random.NewSecureString(32)
	loggedInUsers[token] = user
	cookie := &http.Cookie{Name: "Auth", Value: token, Path: "/", Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}

// Remove .
func RemoveSession(token string) {
	_, ok := loggedInUsers[token]
	if ok {
		delete(loggedInUsers, token)
	}
}

// GetSession .
func GetSession(token string) (*User, bool) {
	user, ok := loggedInUsers[token]
	return user, ok
}

// IsCodeValid checks to see if code provided matches code sent by email
func IsCodeValid(code, email string) bool {
	if codes[email] == code {
		delete(codes, email)
		return true
	}
	return false
}

// GenerateCode creates a code and stores it for later validation
func generateCode(email string) string {
	code := random.NewString(5)
	codes[email] = code
	fmt.Println(code)
	return code
}

// Authenticate .
func Authenticate(w http.ResponseWriter, r *http.Request, c *router.Context) bool {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		return false
	}
	u, ok := GetSession(cookie.Value)
	if ok {
		c.Keys["user"] = u
		return true
	}
	fmt.Println("not authorized")
	w.WriteHeader(http.StatusUnauthorized)
	return false
}
