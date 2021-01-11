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
		fmt.Println(credentials)
		if IsPasswordCorrect(credentials) {
			account, err := getAccount(credentials.Email)
			fmt.Println(account)
			cookie := session.Add(account)

			fmt.Println("user logged in")
			if err != nil {
				fmt.Println("Failed to get user")
			}
			http.SetCookie(w, &cookie)
			fmt.Println(account.Username)
			json.NewEncoder(w).Encode(account)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	router.AuthPost("/login", func(w http.ResponseWriter, r *http.Request, a *account.Account) {

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
		serverStr := []byte(r.FormValue("server"))
		var server server.Server
		json.Unmarshal(serverStr, &server)
		server.ImageURL = util.SaveImage(r)
		var args []interface{}
		args = append(args, server.Name, server.ImageURL)
		fmt.Println(server.Name)
		serverID, err := database.Exec("INSERT INTO Servers (Name, ImageURL) Values (?, ?);", args)
		if err != nil {
			fmt.Println("failed to add server")
		}
		a.AddConnection(serverID, permissions.Full)
	})

	router.AuthPost("/join-server", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		resp, _ := ioutil.ReadAll(r.Body)
		var invite server.Invite
		if err := json.Unmarshal(resp, &invite); err != nil {
			panic(err)
		}
		fmt.Println(invite.Code)
		var args []interface{}
		args = append(args, invite.Code)
		rows, err := database.Query("SELECT ServerID FROM Invites WHERE Code=?;", args)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if rows.Next() {
			var serverID int
			rows.Scan(&serverID)
			a.AddConnection(serverID, permissions.None)
		}

	})
	//	Create TABLE Posts (MessageID Int NOT NULL AUTO_INCREMENT, ServerID Int NOT NULL, UserID Int NOT NULL, Text VARCHAR(255), MediaURL VARCHAR(255), PostDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP, Primary Key (MessageID))

	router.AuthPost("/post", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		resp, _ := ioutil.ReadAll(r.Body)
		var post server.Post
		if err := json.Unmarshal(resp, &post); err != nil {
			panic(err)
		}
		session.Post(a, post)
		var args []interface{}
		args = append(args, post.ServerID, a.ID, post.Text, post.MediaURL)
		_, err := database.Exec("INSERT INTO Posts (ServerID, UserID, Text, MediaURL) Values (?, ?, ?, ?);", args)
		if err != nil {
			panic(err.Error())
		}
	})

	router.AuthGet("/posts", func(w http.ResponseWriter, r *http.Request, a *account.Account) {
		queryValues := r.URL.Query()
		//fmt.Println(queryValues.ServerID)
		var args []interface{}
		args = append(args, queryValues)
		rows, err := database.Query("SELECT ServerID, UserID, MediaURL, Text, TimePosted FROM Posts WHERE ServerID=?", args)
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
