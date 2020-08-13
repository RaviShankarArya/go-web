package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/first.gohtml"))
}

func main() {
	p1 := person{
		Name: "Ravi Shankar",
		Age:  30,
	}

	p2 := person{
		Name: "Sandhya Murthy",
		Age:  27,
	}

	p3 := person{
		Name: "Rushant",
		Age:  6,
	}

	p4 := person{
		Name: "Roopa",
		Age:  33,
	}

	p5 := person{
		Name: "Sushant",
		Age:  34,
	}
	names := []person{p1, p2, p3, p4, p5}
	err := tpl.Execute(os.Stdout, names)

	if err != nil {
		log.Fatalln(err)
	}
}
