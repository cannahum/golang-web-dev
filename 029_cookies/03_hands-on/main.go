package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const cookieName = "website-visit-counter"

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(foo))
}

func foo(w http.ResponseWriter, req *http.Request) {
	var count int
	c1, err := req.Cookie(cookieName)
	if err != nil {
		log.Println("Could not read cookie", err.Error())
	} else {
		c, err := strconv.ParseInt(c1.Value, 10, 32)
		if err != nil {
			log.Println("Could not parse int", err.Error())
		} else {
			count = int(c)
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: fmt.Sprintf("%d", count+1),
	})

	fmt.Fprintf(w, "You have visited %d many times\n", count)
	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}
