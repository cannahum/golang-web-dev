package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type Header struct {
	Title string
}

type NavItem struct {
	Route string
	Name  string
}

type Navigation []NavItem
type PageData struct {
	Header Header
	Nav    Navigation
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("favicon asked.")
		http.ServeFile(w, r, "./public/favicon.ico")
	}))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", PageData{
		Header: Header{
			Title: "JN - Home",
		},
		Nav: []NavItem{
			NavItem{
				Route: "/about",
				Name:  "About",
			},
			NavItem{
				Route: "/contact",
				Name:  "Contact Us",
			},
		},
	})
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", PageData{
		Header: Header{
			Title: "JN - About",
		},
		Nav: []NavItem{
			NavItem{
				Route: "/",
				Name:  "Home",
			},
			NavItem{
				Route: "/contact",
				Name:  "Contact Us",
			},
		},
	})
}

func contact(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact.gohtml", PageData{
		Header: Header{
			Title: "JN - Contact Us",
		},
		Nav: []NavItem{
			NavItem{
				Route: "/",
				Name:  "Home",
			},
			NavItem{
				Route: "/about",
				Name:  "About",
			},
		},
	})
}
