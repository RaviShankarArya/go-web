package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":3000", nil)
}

func setCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "my-cookie", Value: "Welcome to Sanvi Technologies"})
	fmt.Fprintln(w, "COOKIE WRITTEN IN BROWSER")
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Your Cookie:", cookie)
}
