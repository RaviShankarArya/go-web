package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template
var dbUser = map[string]user{}
var dbSession = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{Name: "session", Value: sID.String()}
		http.SetCookie(w, cookie)
	}

	var u user

	if un, ok := dbSession[cookie.Value]; ok {
		u = dbUser[un]
	}

	if req.Method == http.MethodPost {
		firstName := req.FormValue("firstName")
		lastName := req.FormValue("lastName")
		userName := req.FormValue("userName")

		u = user{userName, firstName, lastName}
		dbSession[cookie.Value] = userName
		dbUser[userName] = u
		fmt.Println("dbSession:", dbSession)
		fmt.Println("dbUser:", dbUser)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func show(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSession[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUser[un]
	fmt.Println("dbSession:", dbSession)
	fmt.Println("dbUser:", dbUser)
	tpl.ExecuteTemplate(w, "show.gohtml", u)
}
