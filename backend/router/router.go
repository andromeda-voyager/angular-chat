package router

import (
	"fmt"
	"nebula/session"
	"nebula/user"
	"net/http"
)

type route struct {
	method   string
	callback routeFunction
}

type authRoute struct {
	callback authRouteFunction
}

var routes = make(map[string]map[string]route)
var authRoutes = make(map[string]map[string]authRoute)

type routeFunction func(w http.ResponseWriter, r *http.Request)
type authRouteFunction func(w http.ResponseWriter, r *http.Request, a *user.User)

func init() {
	authRoutes["POST"] = make(map[string]authRoute)
	authRoutes["GET"] = make(map[string]authRoute)
	authRoutes["DELETE"] = make(map[string]authRoute)
	routes["POST"] = make(map[string]route)
	routes["GET"] = make(map[string]route)
	routes["DELETE"] = make(map[string]route)
}

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
func Post(path string, callback routeFunction) {
	routes["POST"][path] = route{
		callback: callback,
	}
}

// Delete registers a callback function for the provided path
func Delete(path string, callback routeFunction) {
	routes["DELETE"][path] = route{
		callback: callback,
	}
}

// Get registers a callback function for the provided path
func Get(path string, callback routeFunction) {
	routes["GET"][path] = route{
		callback: callback,
	}
}

// Post registers a callback function for the provided path
func AuthPost(path string, callback authRouteFunction) {
	authRoutes["POST"][path] = authRoute{
		callback: callback,
	}
}

// Delete registers a callback function for the provided path
func AuthDelete(path string, callback authRouteFunction) {
	authRoutes["DELETE"][path] = authRoute{
		callback: callback,
	}
}

// Get registers a callback function for the provided path
func AuthGet(path string, callback authRouteFunction) {
	authRoutes["GET"][path] = authRoute{
		callback: callback,
	}
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

		fmt.Println(r.URL.String())

		u, ok := authenticate(r)
		if ok {
			callAuthRoute(w, r, u)
		} else {
			callRoute(w, r)
		}
	}
}

func callAuthRoute(w http.ResponseWriter, r *http.Request, u *user.User) {
	authRoute, ok := authRoutes[r.Method][r.URL.String()]
	if ok {
		authRoute.callback(w, r, u)
	} else {
		callRoute(w, r)
	}
}

func callRoute(w http.ResponseWriter, r *http.Request) {
	route, ok := routes[r.Method][r.URL.String()]
	if ok {
		route.callback(w, r)
	}
}
