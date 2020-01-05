package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Listening on port 8080")
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	pageName := request(conn)
	switch pageName {
	case "home":
		respondHome(conn)
	case "about":
		respondAbout(conn)
	default:
		respond404(conn)
	}
}

func request(conn net.Conn) string {
	i := 0
	scanner := bufio.NewScanner(conn)
	var route string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			words := strings.Fields(ln)
			uri := words[1]
			if uri == "/" {
				route = "home"
			} else {
				route = uri[1:]
			}
			break
		}

		i++
	}

	return route
}

func respond404(conn net.Conn) {
	var s bytes.Buffer
	err := tpl.ExecuteTemplate(&s, "404.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	body := s.String()
	fmt.Fprintf(conn, "HTTP/1.1 404\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respondHome(conn net.Conn) {
	var s bytes.Buffer
	err := tpl.ExecuteTemplate(&s, "home.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	body := s.String()
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respondAbout(conn net.Conn) {
	var s bytes.Buffer
	err := tpl.ExecuteTemplate(&s, "about.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	body := s.String()
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
