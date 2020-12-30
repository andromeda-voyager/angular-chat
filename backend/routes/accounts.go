package routes

import (
	"encoding/json"
	"fmt"
	"nebula/accounts"
	"nebula/session"
	"nebula/util"
	"net/http"
)

func init() {

	post("/create-account", func(w http.ResponseWriter, r *http.Request) {
		userStr := []byte(r.FormValue("user"))
		var user accounts.User
		json.Unmarshal(userStr, &user)
		user.AvatarURL = util.SaveImage(r)
		fmt.Println(user)
		if session.IsCodeValid(user.Code, user.Email) {
			fmt.Println("code valid")
			//accounts.Add(user)
		} else {
			fmt.Println("code invalid")
		}

	})

}
