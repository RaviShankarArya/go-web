package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">Set a Cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "mycookie", Value: "Ravi Shankar"})
	fmt.Fprintln(w, `<h1><a href="/read">Read a Cookie</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("mycookie")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
	}
	fmt.Fprintf(w, `<h1>Your Cookie: %v</h1><br/><h1><a href="/expire">Expire</a></h1>`, cookie)
}

func expire(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("mycookie")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
