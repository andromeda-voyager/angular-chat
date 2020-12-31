package accounts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/router"
	"nebula/util"
	"net/http"
)

func init() {

	router.Post("/create-account", func(w http.ResponseWriter, r *http.Request) {
		userStr := []byte(r.FormValue("user"))
		var user User
		json.Unmarshal(userStr, &user)
		user.AvatarURL = util.SaveImage(r)
		fmt.Println(user)
		if isCodeValid(user.Code, user.Email) {
			fmt.Println("added user")
			addUser(user)
			user.Password = ""
			json.NewEncoder(w).Encode(user)
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
			cookie := addSession(credentials.Email)
			fmt.Println("user logged in")
			user, err := getUser(credentials.Email)
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
		var user User
		if err := json.Unmarshal(resp, &user); err != nil {
			panic(err)
		}
		SendCodeToEmail(user.Email)
	})

}
