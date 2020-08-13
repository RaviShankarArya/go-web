package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog_image", dog)
	http.HandleFunc("/image", copy)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<html><body><img src="/dogy.jpeg"></body></html>`)
}

func copy(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dogy.jpeg")
}
