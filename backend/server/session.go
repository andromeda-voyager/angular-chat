package server

import (
	"fmt"

	"golang.org/x/net/websocket"
)

//map(key server id) of map(key userID) of *connection
var servers map[int]map[int]*connection

//map(key channeld id) of map(key userID) of *connection
var channels map[int]map[int]*connection

var connections map[int]*connection

type connection struct {
	ws        *websocket.Conn
	userID    int
	serverID  int
	channelID int
}

func init() {
	servers = make(map[int]map[int]*connection)
	channels = make(map[int]map[int]*connection)
	connections = make(map[int]*connection)
}

func ConnectToChannel(channelID, userID int) {
	c, ok := connections[userID]
	if ok {
		if c.channelID != 0 {
			disconnectFromChannel(c.channelID, userID)
		}
		c.channelID = channelID
		_, ok := channels[channelID]
		if !ok {
			channels[channelID] = make(map[int]*connection)
		}
		channels[channelID][userID] = c
	}
}

func ConnectToServer(serverID, userID int) {
	c, ok := connections[userID]
	if ok {
		if c.serverID != 0 {
			if c.channelID != 0 {
				disconnectFromChannel(c.channelID, userID)
			}
			disconnectFromServer(c.serverID, c.channelID, userID)
		}
		_, ok := servers[c.serverID]
		if !ok {
			servers[c.serverID] = make(map[int]*connection)
		}
		servers[c.serverID][userID] = c
	}
}

func disconnectFromServer(serverID, channelID, userID int) {
	server, ok := servers[serverID]
	if ok {
		_, ok := server[userID]
		if ok {
			delete(server, userID)
		}
	}
}

func disconnectFromChannel(channelID, userID int) {
	channel, ok := channels[channelID]
	if ok {
		_, ok := channel[userID]
		if ok {
			delete(channel, userID)
		}
	}
}

func AddConnection(userID int, ws *websocket.Conn) {
	connections[userID] = &connection{userID: userID, ws: ws, serverID: 0, channelID: 0}
}

func SendMessage(userID int, message string) {
	c, ok := connections[userID]
	if ok {
		if err := websocket.JSON.Send(c.ws, "LO"); err != nil {
			fmt.Println("ws message sent to client")
		}
	}
}

func SendChannelUpdate(u interface{}, channelID int) {
	connections, ok := channels[channelID]
	if ok {
		for _, c := range connections {
			if err := websocket.JSON.Send(c.ws, u); err != nil {
				fmt.Println("ws message sent to client")
				// TODO remove connection
			}
		}
	}
}
