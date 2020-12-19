package main

import (
	"fmt"
	"net/http"
)

type route struct {
	method   string
	callback routeFunction
}

var routes = make(map[string]route)

type routeFunction func(w http.ResponseWriter, r *http.Request)

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Headers", "withcredentials, Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func post(path string, callback routeFunction) {
	routes[path] = route{
		method:   "POST",
		callback: callback,
	}
}

func router(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		setHeaders(&w)
	} else {
		fmt.Println(r.URL.String())
		setHeaders(&w)
		route, ok := routes[r.URL.String()]
		if ok {
			route.callback(w, r)
		}
	}
}
