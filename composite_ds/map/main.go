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
	names := map[string]string{
		"India":          "Delhi",
		"United States":  "Washington",
		"United Kingdom": "London",
		"France":         "Paris",
	}
	err := tpl.Execute(os.Stdout, names)

	if err != nil {
		log.Fatalln(err)
	}
}
