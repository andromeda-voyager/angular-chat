package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/permissions"
	"nebula/router"
	"nebula/util"
	"net/http"
)

func init() {

	router.Post("/create-account", func(w http.ResponseWriter, r *http.Request) {
		accountStr := []byte(r.FormValue("user"))
		var a Account
		json.Unmarshal(accountStr, &a)
		a.AvatarURL = util.SaveImage(r)
		fmt.Println(a)
		if isCodeValid(a.Code, a.Email) {
			fmt.Println("added user")
			addAccount(a)
			a.Password = ""
			json.NewEncoder(w).Encode(a)
		} else {
			fmt.Println("code invalid")
		}
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var credentials Credentials
		if err := json.Unmarshal(resp, &credentials); err != nil {
			panic(err)
		}
		if IsPasswordCorrect(credentials) {
			user, err := getUser(credentials.Email)
			cookie := addSession(user)
			fmt.Println("user logged in")
			if err != nil {
				fmt.Println("Failed to get user")
			}
			json.NewEncoder(w).Encode(user)
			http.SetCookie(w, &cookie)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	router.Post("/send-verification-code", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var a Account
		if err := json.Unmarshal(resp, &a); err != nil {
			panic(err)
		}
		SendCodeToEmail(a.Email)
	})

	router.Post("/create-server", func(w http.ResponseWriter, r *http.Request) {
		serverStr := []byte(r.FormValue("server"))
		var server Server
		json.Unmarshal(serverStr, &server)
		server.ServerImageURL = util.SaveImage(r)
		serverID, err := addServer(server)
		if err != nil {
			fmt.Println("failed to add server")
		}

		cookie, err := r.Cookie("Auth")
		if err != nil {
			fmt.Println("failed to get session")
		}
		a := GetSession(cookie.Value)
		a.AddConnection(serverID, permissions.Full)
	})

}
