package main

import (
	"io"
	"net/http"
)

type hotDog int
type hotCat int

func (d hotDog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Doggy Doggy")
}

func (c hotCat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Kitty Kitty")
}

func main() {
	var d hotDog
	var c hotCat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
