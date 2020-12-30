package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/accounts"
	"nebula/session"
	"net/http"
)

func init() {

	post("/login", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var credentials accounts.Credentials
		if err := json.Unmarshal(resp, &credentials); err != nil {
			panic(err)
		}
		if accounts.IsPasswordCorrect(credentials) {
			cookie := session.Add(credentials.Email)
			fmt.Println("user logged in")
			//	json.NewEncoder(w).Encode(accounts.Get(credentials.Email))
			http.SetCookie(w, &cookie)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	post("/send-verification-code", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var user accounts.User
		if err := json.Unmarshal(resp, &user); err != nil {
			panic(err)
		}
		session.SendCodeToEmail(user.Email)
	})
}
