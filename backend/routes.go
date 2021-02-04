package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/database"
	"nebula/permissions"
	"nebula/router"
	"nebula/server"
	"nebula/session"
	"nebula/user"
	"nebula/util"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

// LoginResponse .
type LoginResponse struct {
	User    *user.User       `json:"user"`
	Servers []*server.Server `json:"servers"`
}

// ServerRequest .
type ServerRequest struct {
	ServerID  int            `json:"serverID"`
	ChannelID int            `json:"channelID"`
	Channel   server.Channel `json:"channel"`
	Role      server.Role    `json:"role"`
	Roles     []server.Role  `json:"roles"`
}

func socket(ws *websocket.Conn) {
	cookie, err := ws.Request().Cookie("Auth")
	if err != nil {
		defer ws.Close()
	} else {
		user, _ := session.Get((cookie.Value))
		servers := getServers(user.ID)
		loginResponse := &LoginResponse{User: user, Servers: servers}
		if err := websocket.JSON.Send(ws, loginResponse); err != nil {
			fmt.Println("unable to send")
			defer ws.Close()
		}
	}
	// 	var m server.Message
	// 	websocket.JSON.Receive(ws, &m)

}

func init() {

	router.AuthPost("/ws", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		s := websocket.Handler(socket)
		s.ServeHTTP(w, r)
	})

	router.Post("/account", func(w http.ResponseWriter, r *http.Request) {
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
		user, err := user.Get(credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			cookie := session.Add(user)
			http.SetCookie(w, cookie)
			json.NewEncoder(w).Encode(user)
		}
	})

	router.AuthPost("/logout", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		cookie, err := r.Cookie("Auth")
		if err == nil {
			session.Remove(cookie.Value)
		}
	})

	router.AuthGet("/login", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		json.NewEncoder(w).Encode(u)
	})

	router.Post("/send-verification-code", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var a user.Account
		if err := json.Unmarshal(resp, &a); err != nil {
			panic(err)
		}
		user.SendCodeToEmail(a.Email)
	})

	router.AuthGet("/server", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		servers := getServers(u.ID)
		json.NewEncoder(w).Encode(servers)
	})

	router.AuthPost("/server", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		serverJSON := []byte(r.FormValue("server"))
		fmt.Println(serverJSON)
		var serverOwner = &server.Member{AccountID: u.ID, Alias: u.Username, Avatar: u.Avatar}
		server := server.New(serverOwner, r)
		json.NewEncoder(w).Encode(server)
	})

	router.AuthPost("/channel", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		resp, _ := ioutil.ReadAll(r.Body)
		var sr *ServerRequest
		if err := json.Unmarshal(resp, &sr); err != nil {
			panic(err)
		}
		if u.HasPermission(permissions.ManageChannels, sr.ServerID) {
			channel := server.NewChannel(sr.Channel, sr.Roles, sr.ServerID)
			json.NewEncoder(w).Encode(channel)
		} else {
			fmt.Println("can't create channel")
		}
	})

	router.AuthDelete("/server", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		serverID, err := strconv.Atoi(r.URL.Query().Get("serverID"))
		if err != nil {
		}
		if u.HasPermission(permissions.Full, serverID) {
			ok := server.Delete(serverID)
			fmt.Println("server deleted?", ok)
		} else {
			fmt.Println("no delete permissions")
		}
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

	router.AuthPost("/message", func(w http.ResponseWriter, r *http.Request, u *user.User) {
		resp, _ := ioutil.ReadAll(r.Body)
		var m server.Message
		if err := json.Unmarshal(resp, &m); err != nil {
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
