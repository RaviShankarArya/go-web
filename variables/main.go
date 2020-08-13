package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/first.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, "Ravi Shankar")

	if err != nil {
		log.Fatalln(err)
	}
}
