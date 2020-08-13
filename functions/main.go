package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("first.gohtml"))
}

func main() {
	p1 := person{
		Name: "Ravi Shankar",
		Age:  30,
	}

	p2 := person{
		Name: "Sandhya",
		Age:  27,
	}

	names := []person{p1, p2}

	data := struct {
		Employees []person
	}{
		names,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "first.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}
}
