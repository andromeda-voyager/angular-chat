package user

import (
	"nebula/router"
	"net/http"
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
