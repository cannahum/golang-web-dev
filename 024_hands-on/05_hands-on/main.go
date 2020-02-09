package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("templates/index.gohtml")
		if err != nil {
			log.Fatalln(err)
		}
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
