package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Your serach key :"+v)
}
