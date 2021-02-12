package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nebula/images"
	"nebula/router"
	"net/http"
)

const DefaultImageUrl = "default-avatar.jpg"

func init() {

	g := router.NewGroup()
	authGroup := router.NewGroup()
	authGroup.Use(Authenticate)

	g.Post("/accounts", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		accountForm := []byte(r.FormValue("user"))
		var a Account
		if err := json.Unmarshal(accountForm, &a); err != nil {
			panic(err)
		}
		a.Avatar = images.Save(r, DefaultImageUrl)
		if IsCodeValid(a.Code, a.Email) {
			u := Add(a)
			cookie := AddSession(u)
			http.SetCookie(w, cookie)
			json.NewEncoder(w).Encode(u)
		} else {
			fmt.Println("code invalid")
		}
	})

	g.Post("/accounts/login", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		resp, _ := ioutil.ReadAll(r.Body)
		var credentials Credentials
		if err := json.Unmarshal(resp, &credentials); err != nil {
			panic(err)
		}
		u, err := Get(credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			cookie := AddSession(u)
			http.SetCookie(w, cookie)
			json.NewEncoder(w).Encode(u)
		}
	})

	g.Post("/accounts/logout", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		cookie, err := r.Cookie("Auth")
		if err == nil {
			RemoveSession(cookie.Value)
		}
	})

	authGroup.Get("/accounts/login", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		u := c.Keys["user"].(*User)
		json.NewEncoder(w).Encode(u)
	})

	g.Post("/accounts/send-verification-code", func(w http.ResponseWriter, r *http.Request, c *router.Context) {
		resp, _ := ioutil.ReadAll(r.Body)
		var a Account
		if err := json.Unmarshal(resp, &a); err != nil {
			panic(err)
		}
		SendCodeToEmail(a.Email)
	})

}
