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

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Name   string
	Hotels []hotel
}

func main() {
	r1 := region{
		Name: "Southern",
		Hotels: []hotel{
			{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			{
				Name:    "Hotel California_1",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
		},
	}

	r2 := region{
		Name: "Northern",
		Hotels: []hotel{
			{
				Name:    "Hotel California_2",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			{
				Name:    "Hotel California_3",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
		},
	}

	r3 := region{
		Name: "Central",
		Hotels: []hotel{
			{
				Name:    "Hotel California_4",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			{
				Name:    "Hotel California_5",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
		},
	}

	hotelByRegion := []region{r1, r2, r3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotelByRegion)

	if err != nil {
		log.Fatalln(err)
	}

}
