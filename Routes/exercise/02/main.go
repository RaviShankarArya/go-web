package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Welocome to index Page")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Doggy")
}

func myName(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", "Hello Ravi SHankar G")

	if err != nil {
		log.Fatalln(err)
	}
}
