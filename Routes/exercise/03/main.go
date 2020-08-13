package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type home int
type dog int
type myName int

func main() {
	var h home
	var d dog
	var m myName
	http.Handle("/", h)
	http.Handle("/dog/", d)
	http.Handle("/me/", m)
	http.ListenAndServe(":3000", nil)
}

func (h home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Welocome to index Page")
}

func (d dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Doggy")
}

func (m myName) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", "Hello Ravi SHankar G")

	if err != nil {
		log.Fatalln(err)
	}
}
