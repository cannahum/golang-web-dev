package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

const (
	AboutPageCookie   = "about-page-visit"
	ContactPageCookie = "contact-page-visit"
	HomePageCookie    = "home-page-cookie"
)

type Header struct {
	Title string
}

type NavItem struct {
	Route string
	Name  string
}

type Navigation []NavItem
type PageData struct {
	Header     Header
	Nav        Navigation
	VisitStats map[string]int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("favicon asked.")
		http.ServeFile(w, r, "./public/favicon.ico")
	}))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(HomePageCookie)
	if err != nil {
		fmt.Println("[index] Error getting cookie", err.Error())
		c = &http.Cookie{
			Name:  HomePageCookie,
			Value: "0",
			Path:  "/",
		}
	}

	visitCount, _ := strconv.Atoi(c.Value)
	c.Value = strconv.Itoa(visitCount + 1)
	http.SetCookie(w, c)

	stats := getCookieAndVisitStats(r)

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
		VisitStats: stats,
	})
}

func about(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(AboutPageCookie)
	if err != nil {
		fmt.Println("[index] Error getting cookie", err.Error())
		c = &http.Cookie{
			Name:  AboutPageCookie,
			Value: "0",
			Path:  "/about",
		}
	}

	visitCount, _ := strconv.Atoi(c.Value)
	c.Value = strconv.Itoa(visitCount + 1)
	http.SetCookie(w, c)

	stats := getCookieAndVisitStats(r)

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
		VisitStats: stats,
	})
}

func contact(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(ContactPageCookie)
	if err != nil {
		fmt.Println("[index] Error getting cookie", err.Error())
		c = &http.Cookie{
			Name:  ContactPageCookie,
			Value: "0",
			Path:  "/contact",
		}
	}

	visitCount, _ := strconv.Atoi(c.Value)
	c.Value = strconv.Itoa(visitCount + 1)
	http.SetCookie(w, c)

	stats := getCookieAndVisitStats(r)

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
		VisitStats: stats,
	})
}

func getCookieAndVisitStats(r *http.Request) map[string]int {
	var cookieMap map[string]int = map[string]int{}

	homeCookie, err := r.Cookie(HomePageCookie)
	if err != nil {
		homeCookie = &http.Cookie{
			Value: "0",
		}
	}
	homePageVisit, _ := strconv.Atoi(homeCookie.Value)
	cookieMap[HomePageCookie] = homePageVisit

	aboutCookie, err := r.Cookie(AboutPageCookie)
	if err != nil {
		aboutCookie = &http.Cookie{
			Value: "0",
		}
	}
	aboutPageVisit, _ := strconv.Atoi(aboutCookie.Value)
	cookieMap[AboutPageCookie] = aboutPageVisit

	contactCookie, err := r.Cookie(ContactPageCookie)
	if err != nil {
		contactCookie = &http.Cookie{
			Value: "0",
		}
	}
	contactPageVisit, _ := strconv.Atoi(contactCookie.Value)
	cookieMap[ContactPageCookie] = contactPageVisit

	return cookieMap
}
