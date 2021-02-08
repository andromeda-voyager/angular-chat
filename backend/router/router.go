package router

import (
	"fmt"
	"nebula/session"
	"nebula/user"
	"net/http"
)

var routes *route

func init() {
	routes = &route{Name: "/", nestedRoutes: make(map[string]*route)}
}

type routeFunction func(w http.ResponseWriter, r *http.Request, c *Context)

func setHeaders(w *http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.String())
	// if strings.HasPrefix(r.URL.String(), "/avatars/") {
	// 	(*w).Header().Set("Content-Type", "image/jpeg")
	// 	fmt.Println("1")
	// } else {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Headers", "withCredentials, Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

// Post registers a callback function for the provided path
func Post(path string, requiresAuth bool, callback routeFunction) {
	routes.Build(splitPath(path, "POST"), callback)
}

// Delete registers a callback function for the provided path
func Delete(path string, requiresAuth bool, callback routeFunction) {
	routes.Build(splitPath(path, "DELETE"), callback)
}

// Get registers a callback function for the provided path
func Get(path string, requiresAuth bool, callback routeFunction) {
	routes.Build(splitPath(path, "GET"), callback)
}

func authenticate(r *http.Request) (*user.User, bool) {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		fmt.Println("no cookie")
		return nil, false
	}
	return session.Get(cookie.Value)
}

// Handler is the router handler function to route paths to functions registered with the router
func Handler(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w, r)
	if r.Method != "OPTIONS" {
		fmt.Println(r.URL.Path)
		c := &Context{Keys: make(map[string]interface{})}
		u, ok := authenticate(r)
		if ok {
			c.Keys["user"] = u
		}
		routes.Match(splitPath(r.URL.Path, r.Method), w, r, c)
	}
}
