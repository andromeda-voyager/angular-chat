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
		session.AddConnection(user.ID, ws)
		// if err := websocket.JSON.Send(ws, loginResponse); err != nil {
		// 	fmt.Println("unable to send")
		// 	defer ws.Close()
		// }
	}

}

func init() {

	router.Post("/ws", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		s := websocket.Handler(socket)
		s.ServeHTTP(w, r)
	})

	router.Get("/test-websocket", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		session.SendMessage(u.ID, "Lo")
	})

	router.Get("/channel/:id/connect", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		channelID := c.Keys["id"].(int)
		session.ConnectToChannel(channelID, u.ID)
	})

	router.Post("/account", false, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
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
			json.NewEncoder(w).Encode(user)
		} else {
			fmt.Println("code invalid")
		}
	})

	router.Post("/login", false, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		fmt.Println("attempting login")
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

	router.Post("/logout", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		cookie, err := r.Cookie("Auth")
		if err == nil {
			session.Remove(cookie.Value)
		}
	})

	router.Get("/login", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		fmt.Println("attempting cookie login")

		if u, auth := c.Keys["user"].(*user.User); auth {
			json.NewEncoder(w).Encode(u)
		}
	})

	router.Post("/send-verification-code", false, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		resp, _ := ioutil.ReadAll(r.Body)
		var a user.Account
		if err := json.Unmarshal(resp, &a); err != nil {
			panic(err)
		}
		user.SendCodeToEmail(a.Email)
	})

	router.Get("/server/:id<int>/connect", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverID := c.Keys["id"].(int)
		s, ok := server.Get(serverID, u.ID)
		if ok {
			json.NewEncoder(w).Encode(s)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	router.Get("/server", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		//serverID, err := strconv.Atoi(r.URL.Query().Get("serverID"))
		servers := getServers(u.ID)
		json.NewEncoder(w).Encode(servers)
	})

	router.Post("/server", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverJSON := []byte(r.FormValue("server"))
		fmt.Println(serverJSON)
		var serverOwner = &server.Member{AccountID: u.ID, Alias: u.Username, Avatar: u.Avatar}
		server := server.New(serverOwner, r)
		json.NewEncoder(w).Encode(server)
	})

	router.Post("/channel", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		resp, _ := ioutil.ReadAll(r.Body)
		var channel *server.Channel
		if err := json.Unmarshal(resp, &channel); err != nil {
			panic(err)
		}
		if u.HasPermission(permissions.ManageChannels, channel.ServerID) {
			server.NewChannel(channel)
			json.NewEncoder(w).Encode(channel)
		} else {
			fmt.Println("can't create channel")
		}
	})

	router.Delete("/server/:id<int>", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverID := c.Keys["id"].(int)
		//		serverID, err := strconv.Atoi(r.URL.Query().Get("serverID"))
		// if err != nil {
		// }
		if u.HasPermission(permissions.Full, serverID) {
			ok := server.Delete(serverID)
			fmt.Println("server deleted?", ok)
		} else {
			fmt.Println("no delete permissions")
		}
	})

	router.Post("/join-server", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		// resp, _ := ioutil.ReadAll(r.Body)
		// var invite server.Invite
		// if err := json.Unmarshal(resp, &invite); err != nil {
		// 	panic(err)
		// }
		// fmt.Println("invite code:", invite.Code)
		// var args []interface{}
		// args = append(args, invite.Code)
		// rows, err := database.Query("SELECT server_id FROM Invite WHERE code=?;", args)
		// if err != nil {
		// 	panic(err)
		// }
		// defer rows.Close()

		// if rows.Next() {
		// 	var serverID int
		// 	rows.Scan(&serverID)
		// }
	})

	router.Post("/message", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
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

	router.Get("/posts", true, func(w http.ResponseWriter, r *http.Request, c *router.Context) {
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
