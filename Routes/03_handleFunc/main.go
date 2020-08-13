package main

import (
	"io"
	"net/http"
)

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Doggy Doggy")
}

func main() {
	http.HandleFunc("/cat", c)
	http.HandleFunc("/dog/", d)
	http.ListenAndServe(":8080", nil)
}
