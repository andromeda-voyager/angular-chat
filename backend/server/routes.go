package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/permissions"
	"nebula/user"
	"net/http"

	router "github.com/andromeda-voyager/go-router"
)

func init() {

	authGroup := router.NewGroup()
	authGroup.Use(user.Authenticate)

	authGroup.Get("/servers/:id<int>/connect", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(user.User)
		serverID := c.Keys["id"].(int)
		ConnectToServer(serverID, u.ID)
		s, ok := Get(serverID, u.ID)
		if ok {
			json.NewEncoder(w).Encode(s)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	authGroup.Post("/servers/:id<int>/invite", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		// u := c.Keys["user"].(user.User)
		fmt.Println("creating code")
		serverID := c.Keys["id"].(int)
		invite := NewInvite(serverID)
		// if ok {
		json.NewEncoder(w).Encode(invite)
		// } else {
		// 	w.WriteHeader(http.StatusBadRequest)
		// }
	})

	authGroup.Get("/servers", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(user.User)
		//serverID, err := strconv.Atoi(r.URL.Query().Get("serverID"))
		servers := getServers(u.ID)
		json.NewEncoder(w).Encode(servers)
	})

	authGroup.Post("/servers", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(user.User)
		server := New(u, r)
		json.NewEncoder(w).Encode(server)
	})

	authGroup.Post("/channels", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(user.User)
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
		u := c.Keys["user"].(user.User)
		serverID := c.Keys["id"].(int)
		if u.HasPermission(permissions.Full, serverID) {
			ok := Delete(serverID)
			fmt.Println("server deleted?", ok)
		} else {
			fmt.Println("no delete permissions")
		}
	})

	authGroup.Post("/servers/join", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(user.User)
		resp, _ := ioutil.ReadAll(r.Body)
		var invite Invite
		if err := json.Unmarshal(resp, &invite); err != nil {
			panic(err)
		}
		s, ok := Join(u, invite)
		if ok {
			json.NewEncoder(w).Encode(s)
		}
	})

	authGroup.Post("/channels/:id<int>/messages", func(w http.ResponseWriter, r *http.Request, c *router.Context) { //TODO check user permissions
		u := c.Keys["user"].(user.User)
		resp, _ := ioutil.ReadAll(r.Body)
		var m Message
		if err := json.Unmarshal(resp, &m); err != nil {
			panic(err)
		}
		m.Add(u.ID)
		update := MessageUpdate{Type: NEW, Event: MESSAGE, Message: m}
		SendChannelUpdate(update, m.ChannelID)
		json.NewEncoder(w).Encode(m)
	})

	authGroup.Delete("/channels/:cID<int>/messages/:mID<int>", func(w http.ResponseWriter, r *http.Request, c *router.Context) { //TODO check user permissions
		// u := c.Keys["user"].(user.User)
		messageID := c.Keys["mID"].(int)
		channelID := c.Keys["cID"].(int)

		deleteMessage(messageID)
		update := MessageUpdate{Type: DELETE, Event: MESSAGE, Message: Message{ID: messageID}}
		SendChannelUpdate(update, channelID)
		json.NewEncoder(w).Encode("ok")
	})

	authGroup.Get("/channels/:id<int>/connect", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(user.User)
		channelID := c.Keys["id"].(int)
		ConnectToChannel(channelID, u.ID)
		messages := GetMessages(channelID)
		json.NewEncoder(w).Encode(messages)
	})

	authGroup.Get("/channels/:id<int>/messages", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		//u := c.Keys["user"].(user.User)
		channelID := c.Keys["id"].(int)
		messages := GetMessages(channelID)
		json.NewEncoder(w).Encode(messages)
	})

	authGroup.Put("/channels/:cID<int>/messages/id", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		// u := c.Keys["user"].(user.User)
		channelID := c.Keys["cID"].(int)
		// messageID := c.Keys["mID"].(int)
		var m Message
		resp, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(resp, &m); err != nil {
			panic(err)
		}
		editMessage(m)
		update := MessageUpdate{Type: MODIFY, Event: MESSAGE, Message: m}
		SendChannelUpdate(update, channelID)
		json.NewEncoder(w).Encode("ok")
	})

}
