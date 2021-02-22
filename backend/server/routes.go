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

	authGroup.Get("/servers/:id<int>/connect", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverID := c.Keys["id"].(int)
		ConnectToServer(serverID, u.ID)
		s, ok := Get(serverID, u.ID)
		if ok {
			json.NewEncoder(w).Encode(s)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	authGroup.Get("/servers", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		//serverID, err := strconv.Atoi(r.URL.Query().Get("serverID"))
		servers := getServers(u.ID)
		json.NewEncoder(w).Encode(servers)
	})

	authGroup.Post("/servers", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		var serverOwner = &Member{AccountID: u.ID, Alias: u.Username, Avatar: u.Avatar}
		server := New(serverOwner, r)
		json.NewEncoder(w).Encode(server)
	})

	authGroup.Post("/channels", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		resp, _ := ioutil.ReadAll(r.Body)
		var channel *Channel
		if err := json.Unmarshal(resp, &channel); err != nil {
			panic(err)
		}
		if u.HasPermission(permissions.ManageChannels, channel.ServerID) {
			channel.Add()
			json.NewEncoder(w).Encode(channel)
		} else {
			fmt.Println("can't create channel")
		}
	})

	authGroup.Delete("/servers/:id<int>", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		serverID := c.Keys["id"].(int)
		if u.HasPermission(permissions.Full, serverID) {
			ok := Delete(serverID)
			fmt.Println("server deleted?", ok)
		} else {
			fmt.Println("no delete permissions")
		}
	})

	authGroup.Post("/servers/join", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
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

	authGroup.Post("/channels/:id<int>/messages", func(w http.ResponseWriter, r *http.Request, c *router.Context) { //TODO check user permissions
		u := c.Keys["user"].(*user.User)
		resp, _ := ioutil.ReadAll(r.Body)
		var m Message
		if err := json.Unmarshal(resp, &m); err != nil {
			panic(err)
		}
		m.Add(u.ID)
		update := MessageUpdate{Type: NEW, Event: MESSAGE, Message: m}
		SendChannelUpdate(update, u.ID, m.ChannelID)
		json.NewEncoder(w).Encode(m)
	})

	authGroup.Delete("/channels/:cID<int>/messages/:mID<int>", func(w http.ResponseWriter, r *http.Request, c *router.Context) { //TODO check user permissions
		u := c.Keys["user"].(*user.User)
		messageID := c.Keys["mID"].(int)
		channelID := c.Keys["cID"].(int)

		deleteMessage(messageID)
		update := MessageUpdate{Type: DELETE, Event: MESSAGE, Message: Message{ID: messageID}}
		SendChannelUpdate(update, u.ID, channelID)
		json.NewEncoder(w).Encode("ok")
	})

	authGroup.Get("/channels/:id<int>/connect", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*user.User)
		channelID := c.Keys["id"].(int)
		ConnectToChannel(channelID, u.ID)
		messages := GetMessages(channelID)
		json.NewEncoder(w).Encode(messages)
	})

	authGroup.Get("/channels/:id<int>/messages", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		//u := c.Keys["user"].(*user.User)
		channelID := c.Keys["id"].(int)
		messages := GetMessages(channelID)
		json.NewEncoder(w).Encode(messages)
	})

}
