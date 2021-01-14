package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/account"
	"nebula/database"
	"nebula/permissions"
	"nebula/router"
	"nebula/server"
	"nebula/session"
	"nebula/util"
	"net/http"
)

func init() {

	// router.AuthPost("/ws", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
	// 	s := websocket.Handler(socket)
	// 	s.ServeHTTP(w, r)
	// })

	router.Post("/create-account", func(w http.ResponseWriter, r *http.Request) {
		accountStr := []byte(r.FormValue("user"))
		var account account.Account
		json.Unmarshal(accountStr, &account)
		account.AvatarURL = util.SaveImage(r)
		fmt.Println(account)
		if isCodeValid(account.Code, account.Email) {
			fmt.Println("added user")
			addAccount(account)
			account.Password = ""
			cookie := session.Add(&account)
			http.SetCookie(w, &cookie)
			json.NewEncoder(w).Encode(account)
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
			account, err := getAccount(credentials.Email)
			cookie := session.Add(account)
			fmt.Println("user logged in")
			if err != nil {
				fmt.Println("Failed to get user")
			}
			http.SetCookie(w, &cookie)
			json.NewEncoder(w).Encode(account)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	router.AuthPost("/logout", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		cookie, err := r.Cookie("Auth")
		fmt.Println("logout")
		if err != nil {
			fmt.Println("no cookie")
		} else {
			fmt.Println("removing session")
			session.Remove(cookie.Value)
		}
	})

	router.AuthGet("/login-with-cookie", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		json.NewEncoder(w).Encode(a)
	})

	router.Post("/send-verification-code", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var a account.Account
		if err := json.Unmarshal(resp, &a); err != nil {
			panic(err)
		}
		SendCodeToEmail(a.Email)
	})

	router.AuthPost("/create-server", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		server := server.New(r)
		a.CreateConnection(server, permissions.Full)
	})

	router.AuthPost("/delete-server", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		resp, _ := ioutil.ReadAll(r.Body)
		var serverID int
		if err := json.Unmarshal(resp, &serverID); err != nil {
			panic(err)
		}
		ok := a.DeleteServer(serverID)
		fmt.Println("server deleted?", ok)

	})

	router.AuthPost("/join-server", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		resp, _ := ioutil.ReadAll(r.Body)
		var invite server.Invite
		if err := json.Unmarshal(resp, &invite); err != nil {
			panic(err)
		}
		fmt.Println("invite code:", invite.Code)
		var args []interface{}
		args = append(args, invite.Code)
		rows, err := database.Query("SELECT server_id FROM invite WHERE code=?;", args)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if rows.Next() {
			var serverID int
			rows.Scan(&serverID)
			//	a.AddConnection(serverID, permissions.None)
		}

	})
	//	Create TABLE Posts (MessageID Int NOT NULL AUTO_INCREMENT, ServerID Int NOT NULL, UserID Int NOT NULL, Text VARCHAR(255), MediaURL VARCHAR(255), PostDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP, Primary Key (MessageID))

	router.AuthPost("/post", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		resp, _ := ioutil.ReadAll(r.Body)
		var post server.Post
		if err := json.Unmarshal(resp, &post); err != nil {
			panic(err)
		}

		if a.CanPostToServer(post.ServerID) {
			session.Post(a, post)
			var args []interface{}
			args = append(args, post.ServerID, a.ID, post.Text, post.MediaURL)
			_, err := database.Exec("INSERT INTO post (server_id, account_id, text, media) Values (?, ?, ?, ?);", args)
			if err != nil {
				panic(err.Error())
			}
		}
	})

	router.AuthGet("/posts", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		queryValues := r.URL.Query()
		//fmt.Println(queryValues.ServerID)
		var args []interface{}
		args = append(args, queryValues)
		rows, err := database.Query("SELECT server_id, account_id, media, text, time_posted FROM post WHERE server_id=?", args)
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
