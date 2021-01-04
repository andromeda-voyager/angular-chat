package router

import (
	"fmt"
	"nebula/account"
	"nebula/session"
	"net/http"
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

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Headers", "withcredentials, Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
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
	if r.Method == "OPTIONS" {
		setHeaders(&w)
	} else {
		fmt.Println(r.URL.String())
		setHeaders(&w)
		a := authenticate(r)
		if a != nil {
			authRoute, ok := authRoutes[r.URL.String()]
			if ok {
				authRoute.callback(w, r, nil)
			}
		} else {
			route, ok := routes[r.URL.String()]
			if ok {
				route.callback(w, r)
			}

		}
	}
}
