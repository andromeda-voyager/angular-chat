package main

import (
	"fmt"
	"log"
	"nebula/router"
	"net/http"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//database.Reset()

	publicFolder, err := filepath.Abs("./public")
	if err != nil {

	}
	fs := http.FileServer(http.Dir(publicFolder))
	fmt.Println(publicFolder)
	http.HandleFunc("/", router.Handler)

	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func socket(ws *websocket.Conn) {
// 	defer ws.Close()
// 	for {
// 		var m text
// 		if err := websocket.JSON.Receive(ws, &m); err != nil {
// 			fmt.Println("unable to receive")
// 			break
// 		}
// 		m2 := text{"thanks"}
// 		if err := websocket.JSON.Send(ws, m2); err != nil {
// 			fmt.Println("unable to send")
// 			break
// 		}
// 	}
// }
