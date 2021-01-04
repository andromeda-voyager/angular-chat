package main

import (
	"fmt"
	"log"
	"nebula/account"
	"nebula/database"
	"nebula/router"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
)

func main() {

	//testQuery()

	http.Handle("/ws", websocket.Handler(socket))

	http.HandleFunc("/", router.Handler)

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type text struct {
	Text string `json:"text"`
}

func socket(ws *websocket.Conn) {
	for {
		var m text
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			fmt.Println("unable to receive")
			break
		}
		m2 := text{"thanks"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			fmt.Println("unable to send")
			break
		}
	}
}

func testQuery() {

	var args []interface{}
	rows, err := database.Query("SELECT Email FROM Users", args)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var a account.Account
		rows.Scan(&a.Email)
		fmt.Println(a.Email)
	}
}
