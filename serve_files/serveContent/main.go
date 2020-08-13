package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/image", copy)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/dogy.jpeg">`)
}

func copy(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("dogy.jpeg")
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}

	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
