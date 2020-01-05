package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

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
	info := request(conn)
	respond(conn, info)
}

func request(conn net.Conn) string {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()

		if i == 1 {
			words := strings.Fields(ln)
			fmt.Println(words)
			return words[1]
		}

		i++
	}

	return ""
}

func respond(conn net.Conn, info string) {
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html><head><title>Jon Mon</title></head><body><p>Here is your URI %v</p><p>Goodbye</p></body></html>
	`, info)
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
