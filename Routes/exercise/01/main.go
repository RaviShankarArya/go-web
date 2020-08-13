package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Welocome to index Page")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Doggy")
}

func myName(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Ravi Shankar")
}
