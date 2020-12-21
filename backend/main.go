package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"nebula/accounts"
	"nebula/config"
	"nebula/session"
	"nebula/util"
	"net/http"
	"os"
	"path"

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

	post("/create-account", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var user accounts.User
		if err := json.Unmarshal(resp, &user); err != nil {
			panic(err)
		}
		if accounts.DoesAccountExist(user.Email) {
			accounts.Add(user)
		}
	})

	post("/upload-avatar", func(w http.ResponseWriter, r *http.Request) {
		in, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed to get formFile")

		}
		defer in.Close()
		avatarImgPath := "./public/avatars/" + util.NewRandomString(10) + ".jpg"
		fmt.Println(path.Join(config.ServerURL, avatarImgPath))
		out, err := os.Create(avatarImgPath) //header.Filename
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed to open")
		}
		defer out.Close()
		io.Copy(out, in)

		// resp, _ := ioutil.ReadAll(r.Body)
		// var user accounts.User
		// if err := json.Unmarshal(resp, &user); err != nil {
		// 	panic(err)
		// }
		// if accounts.DoesAccountExist(user.Email) {
		// 	accounts.Add(user)
		// }
	})

	http.HandleFunc("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
