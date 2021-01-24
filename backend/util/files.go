package util

import (
	"fmt"
	"io"
	"math/rand"
	"nebula/config"
	"net/http"
	"os"
	"path"
)

const avatarFolder = "./public/images/"
const avatarBaseURL = config.ServerURL + "/static/images/"

// SaveImage .
func SaveImage(r *http.Request) string {
	in, _, err := r.FormFile("image")
	var fileName string
	if err != nil {
		fmt.Println("using default profile img")
		fileName = "default-avatar.jpg"
	} else {
		fileName = newRandomString(10) + ".jpg"
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

// NewRandomString returns a pseudo random alphanumeric string
func newRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
