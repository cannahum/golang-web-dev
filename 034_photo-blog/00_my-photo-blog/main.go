package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/", NewEnsureSession(index))
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request, c *http.Cookie) {
	if r.Method == http.MethodPost {
		file, fileHeader, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
			return
		} else if file == nil {
			w.WriteHeader(http.StatusOK)
			return
		}
		defer file.Close()

		// Get file data to create sha1
		extension := strings.Split(fileHeader.Filename, ".")[1]
		hash := sha1.New()
		io.Copy(hash, file)
		newFileName := fmt.Sprintf("%x.%s", hash.Sum(nil), extension)

		// Create new file locally
		workingDirectory, _ := os.Getwd()
		newFilePath := filepath.Join(workingDirectory, "public", "pics", newFileName)

		newFile, err := os.Create(newFilePath)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		defer newFile.Close()
		file.Seek(0, 0)
		io.Copy(newFile, file)

		// Sort the cookie. Append the file name to the user's cookie
		currentValue := c.Value
		if !strings.Contains(currentValue, newFileName) {
			newValue := strings.Join([]string{currentValue, newFileName}, "|")
			c.Value = newValue
			http.SetCookie(w, c)
		}
	}

	values := strings.Split(c.Value, "|")[1:]
	tpl.ExecuteTemplate(w, "index.gohtml", values)
}
