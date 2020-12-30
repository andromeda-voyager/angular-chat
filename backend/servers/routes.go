package servers

import (
	"encoding/json"
	"nebula/router"
	"nebula/util"
	"net/http"
)

func init() {
	router.Post("/create-server", func(w http.ResponseWriter, r *http.Request) {
		serverStr := []byte(r.FormValue("server"))
		var server Server
		json.Unmarshal(serverStr, &server)
		server.ServerImageURL = util.SaveImage(r)
	})
}
