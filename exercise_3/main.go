package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type item struct {
	Name  string
	Price float64
}

type meal struct {
	Name  string
	Items []item
}

type hotel struct {
	Name  string
	Meals []meal
}

func main() {
	m1 := meal{
		Name: "Breakfast",
		Items: []item{
			{Name: "Idli", Price: 20.0},
			{Name: "Dosa", Price: 40.0},
			{Name: "Poori", Price: 35.0},
			{Name: "Buns", Price: 25.0},
		},
	}

	m2 := meal{
		Name: "Lunch",
		Items: []item{
			{Name: "Fried Rice", Price: 120.0},
			{Name: "Mini Meals", Price: 60.0},
			{Name: "Meals", Price: 90.0},
			{Name: "Noodles", Price: 55.0},
		},
	}

	m3 := meal{
		Name: "Dinner",
		Items: []item{
			{Name: "Mini Meals", Price: 60.0},
			{Name: "Meals", Price: 90.0},
			{Name: "Roti-Curry", Price: 65.0},
		},
	}
	menu := []meal{m1, m2, m3}

	h1 := hotel{
		Name:  "Udupi Hotel",
		Meals: menu,
	}

	h2 := hotel{
		Name:  "Manglore Hotel",
		Meals: menu,
	}
	h3 := hotel{
		Name:  "Davanagere Hotel",
		Meals: menu,
	}

	hotels := []hotel{h1, h2, h3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)

	if err != nil {
		log.Fatalln(err)
	}
}
