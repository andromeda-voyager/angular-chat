package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/permissions"
	"nebula/router"
	"nebula/user"
	"net/http"
)

func init() {

	authGroup := router.NewGroup()
	authGroup.Use(user.Authenticate)

	authGroup.Get("/channel/:id/connect", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		channelID := c.Keys["id"].(int)
		ConnectToChannel(channelID, u.ID)
	})

	authGroup.Get("/server/:id<int>/connect", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverID := c.Keys["id"].(int)
		s, ok := Get(serverID, u.ID)
		if ok {
			json.NewEncoder(w).Encode(s)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	authGroup.Get("/server", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		//serverID, err := strconv.Atoi(r.URL.Query().Get("serverID"))
		servers := getServers(u.ID)
		json.NewEncoder(w).Encode(servers)
	})

	authGroup.Post("/server", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverJSON := []byte(r.FormValue("server"))
		fmt.Println(serverJSON)
		var serverOwner = &Member{AccountID: u.ID, Alias: u.Username, Avatar: u.Avatar}
		server := New(serverOwner, r)
		json.NewEncoder(w).Encode(server)
	})

	authGroup.Post("/channel", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		resp, _ := ioutil.ReadAll(r.Body)
		var channel *Channel
		if err := json.Unmarshal(resp, &channel); err != nil {
			panic(err)
		}
		if u.HasPermission(permissions.ManageChannels, channel.ServerID) {
			NewChannel(channel)
			json.NewEncoder(w).Encode(channel)
		} else {
			fmt.Println("can't create channel")
		}
	})

	authGroup.Delete("/server/:id<int>", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverID := c.Keys["id"].(int)
		if u.HasPermission(permissions.Full, serverID) {
			ok := Delete(serverID)
			fmt.Println("server deleted?", ok)
		} else {
			fmt.Println("no delete permissions")
		}
	})

	authGroup.Post("/join-server", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
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

	authGroup.Post("/channel/:id<int>/messages", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		//u := c.Keys["user"].(*user.User)
		resp, _ := ioutil.ReadAll(r.Body)
		var m Message
		if err := json.Unmarshal(resp, &m); err != nil {
			panic(err)
		}
		fmt.Println(m)
		json.NewEncoder(w).Encode(m)
	})

	authGroup.Get("/channel/:id<int>/messages", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		//u := c.Keys["user"].(*user.User)
		channelID := c.Keys["id"].(int)
		messages := GetMessages(channelID)
		json.NewEncoder(w).Encode(messages)
	})

}
