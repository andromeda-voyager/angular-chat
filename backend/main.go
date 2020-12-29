package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"nebula/accounts"
	"nebula/config"
	"nebula/session"
	"nebula/util"
	"net/http"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
)

func saveImage(r *http.Request) string {
	in, _, err := r.FormFile("image")
	var avatarImgPath string
	if err != nil {
		fmt.Println("using default profile img")
		avatarImgPath = "./public/avatars/default-avatar.jpg"
	} else {
		avatarImgPath = "./public/avatars/" + util.NewRandomString(10) + ".jpg"
		out, err := os.Create(avatarImgPath) //header.Filename
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed to open")
			avatarImgPath = "./public/avatars/default-avatar.jpg"
		} else {
			defer out.Close()
			defer in.Close()
			io.Copy(out, in)
		}
	}
	return path.Join(config.ServerURL, avatarImgPath)
}

func main() {

	//	json.NewEncoder(w).Encode(user)
	//cookieR, err := r.Cookie("NUser")

	post("/login", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var credentials accounts.Credentials
		if err := json.Unmarshal(resp, &credentials); err != nil {
			panic(err)
		}
		if accounts.IsPasswordCorrect(credentials) {
			cookie := session.Add(credentials.Email)
			fmt.Println("user logged in")
			http.SetCookie(w, &cookie)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	post("/create-account", func(w http.ResponseWriter, r *http.Request) {
		userStr := []byte(r.FormValue("user"))
		var user accounts.User
		json.Unmarshal(userStr, &user)
		user.AvatarURL = saveImage(r)
		fmt.Println(user)
		if session.IsCodeValid(user.Code, user.Email) {
			fmt.Println("code valid")
			//accounts.Add(user)
		} else {
			fmt.Println("code invalid")
		}

	})

	post("/send-verification-code", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadAll(r.Body)
		var user accounts.User
		if err := json.Unmarshal(resp, &user); err != nil {
			panic(err)
		}
		session.SendCodeToEmail(user.Email)
	})

	post("/upload-image", func(w http.ResponseWriter, r *http.Request) {
		saveImage(r)
	})

	http.Handle("/ws", websocket.Handler(socket))
	http.HandleFunc("/", router)

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
