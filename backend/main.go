package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	router "github.com/andromeda-voyager/go-router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//database.Reset()
	publicFolder, err := filepath.Abs("./public")
	if err != nil {

	}

	fs := http.FileServer(http.Dir(publicFolder))
	//fmt.Println(publicFolder)

	http.HandleFunc("/", router.Handler)

	http.Handle("/static/", http.StripPrefix("/static", fs))
	//http.Handle("/", Auth(router.Handler))

	fmt.Println("Listening on port 8080")
	//	router.PrintRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
