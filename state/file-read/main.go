package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println("Method:", req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		fmt.Println("File:", f)
		fmt.Println("Header:", h)
		fmt.Println("Error:", err)
		if err != nil {
			http.Error(w, "For Form File: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, "For File Read: "+err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form enctype="multipart/form-data" method="post">
	<input type="file" name="q"/>
	<input type="submit"/>
	</form><br/>
	`+s)
}
