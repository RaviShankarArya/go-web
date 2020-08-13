package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html;Charset=utf-8")
	io.WriteString(w, `
		<form method="GET">
			<input type="text" name="q"/>
			<input type="submit" />
		</form>
	`+v)
}
