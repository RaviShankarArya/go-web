package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template
var dbSession = map[string]string{}
var dbUser = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", nil)
}

type user struct {
	UserName  string
	FirstName string
	LastName  string
	Password  string
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		sId, _ := uuid.NewV4()
		cookie = &http.Cookie{Name: "session", Value: sId.String()}
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
		password := req.FormValue("password")

		u = user{userName, firstName, lastName, password}
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

func signup(w http.ResponseWriter, req *http.Request) {
	if isUserLoggedIn(req) {
		fmt.Println("User already Exists")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		firstName := req.FormValue("firstName")
		lastName := req.FormValue("lastName")
		userName := req.FormValue("userName")
		password := req.FormValue("password")

		sId, _ := uuid.NewV4()
		cookie := &http.Cookie{Name: "session", Value: sId.String()}
		http.SetCookie(w, cookie)

		dbSession[cookie.Value] = userName

		u := user{firstName, lastName, userName, password}
		dbUser[userName] = u

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
