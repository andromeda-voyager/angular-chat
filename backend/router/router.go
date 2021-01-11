package router

import (
	"fmt"
	"nebula/account"
	"nebula/session"
	"net/http"
	"strings"
)

type route struct {
	method   string
	callback routeFunction
}

type authRoute struct {
	method   string
	callback authRouteFunction
}

var routes = make(map[string]route)
var authRoutes = make(map[string]authRoute)

type routeFunction func(w http.ResponseWriter, r *http.Request)
type authRouteFunction func(w http.ResponseWriter, r *http.Request, a *account.Account)

func setHeaders(w *http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	if strings.HasPrefix(r.URL.String(), "/avatars/") {
		(*w).Header().Set("Content-Type", "image/jpeg")
		fmt.Println("1")
	} else {
		(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		(*w).Header().Set("Access-Control-Allow-Headers", "withCredentials, Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Println("2")

		(*w).Header().Set("Content-Type", "application/json")

		(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

// Post registers a callback function for the provided path
func Post(path string, callback routeFunction) {
	routes[path] = route{
		method:   "POST",
		callback: callback,
	}
}

// AuthPost registers a callback function for the provided path
func AuthPost(path string, callback authRouteFunction) {
	authRoutes[path] = authRoute{
		method:   "POST",
		callback: callback,
	}
}

// AuthGet registers a callback function for the provided path
func AuthGet(path string, callback authRouteFunction) {
	authRoutes[path] = authRoute{
		method:   "GET",
		callback: callback,
	}
}

func authenticate(r *http.Request) *account.Account {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		fmt.Println("no cookie")
		return nil
	}
	return session.Get(cookie.Value)
}

// Handler is the router handler function to route paths to functions registered with the router
func Handler(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w, r)

	if r.Method != "OPTIONS" {

		fmt.Println(r.URL.String())

		a := authenticate(r)
		if a != nil {
			authRoute, ok := authRoutes[r.URL.String()]
			if ok {
				authRoute.callback(w, r, a)
			}
		} else {
			route, ok := routes[r.URL.String()]
			if ok {
				route.callback(w, r)
			}

		}
	}
}
