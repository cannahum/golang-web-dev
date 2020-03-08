package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

func index(w http.ResponseWriter, req *http.Request, cookieValue string) {
	tpl.ExecuteTemplate(w, "index.gohtml", cookieValue)
}
