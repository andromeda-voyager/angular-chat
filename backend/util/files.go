package util

import (
	"fmt"
	"io"
	"nebula/config"
	"net/http"
	"os"
	"path"
)

func SaveImage(r *http.Request) string {
	in, _, err := r.FormFile("image")
	var avatarImgPath string
	if err != nil {
		fmt.Println("using default profile img")
		avatarImgPath = "./public/avatars/default-avatar.jpg"
	} else {
		avatarImgPath = "./public/avatars/" + NewRandomString(10) + ".jpg"
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
