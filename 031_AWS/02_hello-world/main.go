package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, check out the size of my server, I'm running on AWS.")
}