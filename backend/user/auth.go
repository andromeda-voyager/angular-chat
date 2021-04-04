package user

import (
	"net/http"

	router "github.com/andromeda-voyager/go-router"
)

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
	w.WriteHeader(http.StatusUnauthorized)
	return false
}
