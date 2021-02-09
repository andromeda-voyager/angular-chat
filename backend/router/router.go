package router

import (
	"fmt"
	"net/http"
)

type account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

var routes *route

func init() {
	routes = &route{Name: "/", nestedRoutes: make(map[string]*route)}
}

type routeCallbackFunc func(w http.ResponseWriter, r *http.Request, c *Context)

func setHeaders(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Headers", "withCredentials, Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

// Post registers a callback function for the provided path
func (g *Group) Post(path string, callback routeCallbackFunc) {
	routes.Build(splitPath(path, "POST"), callback, g)
}

// Delete registers a callback function for the provided path
func (g *Group) Delete(path string, callback routeCallbackFunc) {
	routes.Build(splitPath(path, "DELETE"), callback, g)
}

// Get registers a callback function for the provided path
func (g *Group) Get(path string, callback routeCallbackFunc) {
	routes.Build(splitPath(path, "GET"), callback, g)
}

// Handler is the router handler function to route paths to functions registered with the router
func Handler(w http.ResponseWriter, req *http.Request) {
	setHeaders(&w, req)
	if req.Method != "OPTIONS" {
		fmt.Println(req.URL.Path)
		c := &Context{Keys: make(map[string]interface{})}
		routes.Match(splitPath(req.URL.Path, req.Method), w, req, c)
	}
}
