package main

import (
	"fmt"
	"log"
	"nebula/account"
	"nebula/database"
	"nebula/permissions"
	"nebula/router"
	"net/http"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
)

func main() {

	//testQuery()

	fmt.Println(permissions.CanDeleteServer(192))
	publicFolder, err := filepath.Abs("./public")
	if err != nil {

	}
	fs := http.FileServer(http.Dir(publicFolder))
	fmt.Println(publicFolder)
	http.HandleFunc("/", router.Handler)

	http.Handle("/static/", http.StripPrefix("/static", fs))

	//http.HandleFunc("/", router.Handler)

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type text struct {
	Text string `json:"text"`
}

func socket(ws *websocket.Conn) {
	defer ws.Close()
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
