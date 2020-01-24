package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)

	i := 0
	var method string
	var uri string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}

		if i == 0 {
			// First line has the request line
			fields := strings.Fields(ln)
			method = fields[0]
			uri = fields[1]
		}
		i++
	}

	body := fmt.Sprintf("Response!!! This is cool.\nYou made a %s request to %s", method, uri)
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
