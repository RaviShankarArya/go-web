package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdouble":     fdouble,
	"fSquareroot": fSquareroot,
	"fsquare":     fsquare,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("first.gohtml"))
}

func fdouble(x int) int {
	return x + x
}

func fSquareroot(x float64) float64 {
	return math.Sqrt(x)
}

func fsquare(x int) float64 {
	return math.Pow(float64(x), 2)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "first.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
