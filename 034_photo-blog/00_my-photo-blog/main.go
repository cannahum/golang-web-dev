package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/", NewEnsureSession(index))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request, c *http.Cookie) {
	if r.Method == http.MethodPost {
		imageName := r.FormValue("image-name")
		if imageName == "" {
			w.WriteHeader(http.StatusOK)
		} else {
			currentValue := c.Value
			if !strings.Contains(currentValue, imageName) {
				newValue := strings.Join([]string{currentValue, imageName}, "|")
				c.Value = newValue
				http.SetCookie(w, c)
			}
		}
	}

	values := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", values)
}
