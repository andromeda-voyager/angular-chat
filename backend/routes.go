package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/database"
	"nebula/router"
	"nebula/server"
	"nebula/session"
	"nebula/user"
	"nebula/util"
	"net/http"
)

type LoginResponse struct {
	User    *user.User       `json:"user"`
	Servers []*server.Server `json:"servers"`
}

func init() {

	// router.AuthPost("/ws", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
	// 	s := websocket.Handler(socket)
	// 	s.ServeHTTP(w, r)
	// })

	router.Post("/create-account", func(w http.ResponseWriter, r *http.Request) {
		accountStr := []byte(r.FormValue("user"))
		var a user.Account
		json.Unmarshal(accountStr, &a)
		a.Avatar = util.SaveImage(r)
		fmt.Println(a)
		if user.IsCodeValid(a.Code, a.Email) {
			fmt.Println("added user")
			user := user.Add(a)
			cookie := session.Add(user)
			http.SetCookie(w, cookie)
			loginResponse := &LoginResponse{User: user, Servers: nil}
			json.NewEncoder(w).Encode(loginResponse)
		} else {
			fmt.Println("code invalid")
		}
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var credentials user.Credentials
		if err := json.Unmarshal(resp, &credentials); err != nil {
			panic(err)
		}
		user, err := user.Get(credentials.Email)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if user.IsPasswordCorrect(credentials.Password) {
			fmt.Println("user logged in")
			servers := getServers(user.ID)
			loginResponse := &LoginResponse{User: user, Servers: servers}
			cookie := session.Add(user)
			http.SetCookie(w, cookie)
			json.NewEncoder(w).Encode(loginResponse)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	router.AuthPost("/logout", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		cookie, err := r.Cookie("Auth")
		fmt.Println("logout")
		if err != nil {
			fmt.Println("no cookie")
		} else {
			fmt.Println("removing session")
			session.Remove(cookie.Value)
		}
	})

	router.AuthGet("/login-with-cookie", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		servers := getServers(u.ID)
		loginResponse := &LoginResponse{User: u, Servers: servers}
		json.NewEncoder(w).Encode(loginResponse)
	})

	router.Post("/send-verification-code", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var a user.Account
		if err := json.Unmarshal(resp, &a); err != nil {
			panic(err)
		}
		user.SendCodeToEmail(a.Email)
	})

	router.AuthPost("/create-server", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		serverJSON := []byte(r.FormValue("server"))
		fmt.Println(serverJSON)
		server := server.New(u, r)
		json.NewEncoder(w).Encode(server)
	})

	router.AuthPost("/create-channel", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		serverJSON := []byte(r.FormValue("server"))
		fmt.Println(serverJSON)
		server := server.New(u, r)
		json.NewEncoder(w).Encode(server)
	})

	router.AuthPost("/delete-server", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		resp, _ := ioutil.ReadAll(r.Body)
		var serverID int
		if err := json.Unmarshal(resp, &serverID); err != nil {
			panic(err)
		}
		ok := u.DeleteServer(serverID)
		fmt.Println("server deleted?", ok)

	})

	router.AuthPost("/join-server", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		resp, _ := ioutil.ReadAll(r.Body)
		var invite server.Invite
		if err := json.Unmarshal(resp, &invite); err != nil {
			panic(err)
		}
		fmt.Println("invite code:", invite.Code)
		var args []interface{}
		args = append(args, invite.Code)
		rows, err := database.Query("SELECT server_id FROM Invite WHERE code=?;", args)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if rows.Next() {
			var serverID int
			rows.Scan(&serverID)
		}

	})

	router.AuthPost("/post", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		resp, _ := ioutil.ReadAll(r.Body)
		var post server.Post
		if err := json.Unmarshal(resp, &post); err != nil {
			panic(err)
		}

		// ok := a.CanPostOnServer(post.ServerID)
		// if ok {
		// 	session.Post(a, post)
		// 	var args []interface{}
		// 	args = append(args, post.ServerID, a.ID, post.Text, post.MediaURL)
		// 	_, err := database.Exec("INSERT INTO post (server_id, account_id, text, media) Values (?, ?, ?, ?);", args)
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}
		// }
	})

	router.AuthGet("/posts", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		queryValues := r.URL.Query()
		//fmt.Println(queryValues.ServerID)
		var args []interface{}
		args = append(args, queryValues)
		rows, err := database.Query("SELECT server_id, account_id, media, text, time_posted FROM Post WHERE server_id=?", args)
		if err != nil {
			panic(err.Error())
		}
		if rows.Next() {
			var text int
			rows.Scan(&text)
			fmt.Println(text)
		}
	})

}
