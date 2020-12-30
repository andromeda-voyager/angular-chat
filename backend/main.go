package main

import (
	"fmt"
	"log"
	"nebula/accounts"
	"nebula/routes"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
)

func main() {

	//	json.NewEncoder(w).Encode(user)
	//cookieR, err := r.Cookie("NUser")

	// post("/upload-image", func(w http.ResponseWriter, r *http.Request) {
	// 	saveImage(r)
	// })

	accounts.TestQuery()

	http.Handle("/ws", websocket.Handler(socket))
	http.HandleFunc("/", routes.Router)

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type message struct {
	Text string `json:"text"`
}

func socket(ws *websocket.Conn) {
	for {
		var m message
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			fmt.Println("unable to receive")
			break
		}
		m2 := message{"thanks"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			fmt.Println("unable to send")
			break
		}
	}
}
