package main

import (
	"io"
	"net/http"
)

type hotDog int
type hotCat int

func (d hotDog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Welcome Doggy")
}

func (c hotCat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Welcome Kitty")
}

func main() {
	var d hotDog
	var c hotCat

	http.Handle("/cat", c)
	http.Handle("/dog/", d)
	http.ListenAndServe(":8080", nil)
}
