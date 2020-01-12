package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var indexVisit int
var meVisit int
var dogVisit int

var tpl *template.Template



func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "Welcome to the Page")
	})
	http.HandleFunc("/dog", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "Welcome to the Dog Section")
	})
	http.HandleFunc("/me", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "About me")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
