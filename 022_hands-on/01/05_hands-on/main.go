package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

var indexVisit int
var dogVisit int
var meVisit int

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.Handle("/", http.HandlerFunc(handleIndex))
	http.Handle("/dog/", http.HandlerFunc(handleDog))
	http.Handle("/me/", http.HandlerFunc(handleMe))

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	var s bytes.Buffer
	indexVisit++
	tpl.ExecuteTemplate(&s, "index.gohtml", indexVisit)
	_, _ = fmt.Fprintf(w, s.String())
}

func handleDog(w http.ResponseWriter, req *http.Request) {
	var s bytes.Buffer
	dogVisit++
	tpl.ExecuteTemplate(&s, "dog.gohtml", dogVisit)
	_, _ = fmt.Fprintf(w, s.String())
}

func handleMe(w http.ResponseWriter, req *http.Request) {
	var s bytes.Buffer
	meVisit++
	tpl.ExecuteTemplate(&s, "me.gohtml", meVisit)
	_, _ = fmt.Fprintf(w, s.String())
}
