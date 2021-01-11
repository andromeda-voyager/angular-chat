package util

import (
	"fmt"
	"io"
	"nebula/config"
	"net/http"
	"os"
	"path"
)

const avatarFolder = "./public/avatars/"
const avatarBaseURL = config.ServerURL + "/static/avatars/"

func SaveImage(r *http.Request) string {
	in, _, err := r.FormFile("image")
	var fileName string
	if err != nil {
		fmt.Println("using default profile img")
		fileName = "default-avatar.jpg"
	} else {
		fileName = NewRandomString(10) + ".jpg"
		out, err := os.Create(path.Join(avatarFolder, fileName)) //header.Filename
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed to open")
			fileName = "default-avatar.jpg"
		} else {
			defer out.Close()
			defer in.Close()
			io.Copy(out, in)
		}
	}
	return avatarBaseURL + fileName
}
