package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotDog int

func (h hotDog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Doggy Doggy")
	case "/cat":
		io.WriteString(w, "Cat Cat")
	default:
		fmt.Fprintln(w, "Homepage")
	}
}

func main() {
	var d hotDog
	http.ListenAndServe(":8080", d)
}
