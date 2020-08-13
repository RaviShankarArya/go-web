package main

import (
	"fmt"
	"net/http"
)

type hotDog int

func (h hotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Welcome to Go Web Application")
}

func main() {
	var d hotDog
	http.ListenAndServe(":3000", d)
}
