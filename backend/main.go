package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nebulous/accounts"
	"nebulous/session"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//	json.NewEncoder(w).Encode(user)
	//cookieR, err := r.Cookie("NUser")

	post("/login", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var credentials accounts.Credentials
		if err := json.Unmarshal(resp, &credentials); err != nil {
			panic(err)
		}
		if accounts.IsPasswordCorrect(credentials) {
			cookie := session.Add(credentials.Email)
			fmt.Println("user logged in")
			http.SetCookie(w, &cookie)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	http.HandleFunc("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
