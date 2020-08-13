package main

import "fmt"

func main() {
	name := "Ravi Shankar"
	templ := `<html>
	<head>
		<title>Hello World</title>
	</head>
	<body>
	<h1>` + name + `</h2>
	</body>
	<html>`

	fmt.Println(templ)
}
