package images

import (
	"fmt"
	"io"
	"nebula/config"
	"nebula/random"
	"net/http"
	"os"
	"path"
)

const avatarFolder = "./public/images/"
const avatarBaseURL = config.ServerURL + "/static/images/"

// Save .
func Save(r *http.Request, defaultUrl string) string {
	var fileName string
	imageFile, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("No image provided. Using default image.")
		return defaultUrl
	}
	fileName = random.NewString(10) + ".jpg"
	out, err := os.Create(path.Join(avatarFolder, fileName)) //header.Filename
	if err != nil {
		fmt.Println("Dailed to save image. Using default.")
		return defaultUrl
	} else {
		defer out.Close()
		defer imageFile.Close()
		io.Copy(out, imageFile)
	}

	return avatarBaseURL + fileName
}
