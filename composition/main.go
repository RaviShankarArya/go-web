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

type course struct {
	Name, Number, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{Name: "Introduction to Go", Number: "111", Units: "4"},
				{Name: "Introduction to Ruby", Number: "112", Units: "5"},
				{Name: "Introduction to Python", Number: "113", Units: "6"},
				{Name: "Introduction to Node", Number: "114", Units: "8"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{Name: "Introduction to HTML", Number: "111", Units: "4"},
				{Name: "Introduction to CSS", Number: "112", Units: "5"},
				{Name: "Introduction to Javascript", Number: "113", Units: "6"},
				{Name: "Introduction to React", Number: "114", Units: "8"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				{Name: "Introduction to C#", Number: "111", Units: "4"},
				{Name: "Introduction to Java", Number: "112", Units: "5"},
				{Name: "Introduction to Rails", Number: "113", Units: "6"},
				{Name: "Introduction to Springs", Number: "114", Units: "8"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "first.gohtml", y)

	if err != nil {
		log.Fatalln(err)
	}
}
