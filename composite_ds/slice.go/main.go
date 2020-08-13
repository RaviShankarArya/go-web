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
	names := []string{"Ravi", "Shankar", "Sandhya", "Murthy"}
	err := tpl.ExecuteTemplate(os.Stdout, "first.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}
}
