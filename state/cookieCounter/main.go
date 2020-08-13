package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", counter)
	http.ListenAndServe(":3000", nil)
}

func counter(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie-count")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{Name: "my-cookie-count", Value: "0"}
	}
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "count:", cookie.Value)
}
