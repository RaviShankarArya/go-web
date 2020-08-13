package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotDog int

func (h hotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Fprintln(w, "Hello Welcome to GO!!!")
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}
func main() {
	var d hotDog
	http.ListenAndServe(":8080", d)
}
