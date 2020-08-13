package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("first.gohtml"))
}

type person struct {
	Name string
	Age  int
}

func (p person) Something() int {
	return 7
}

func (p person) Double() int {
	return p.Age * 2
}

func (p person) TakeArgs(x int) int {
	return x * 2
}

func main() {
	p := person{Name: "Ravi", Age: 30}
	err := tpl.ExecuteTemplate(os.Stdout, "first.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
