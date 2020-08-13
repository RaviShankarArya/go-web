package main

import (
	"fmt"
	"net/http"
)

func isUserLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSession[cookie.Value]
	_, ok := dbUser[un]
	return ok
}

func getUser(req *http.Request) user {
	var u user
	cookie, err := req.Cookie("session")
	if err != nil {
		fmt.Println("Err:", err)
		return u
	}
	if un, ok := dbSession[cookie.Value]; ok {
		dbUser[un] = u
	}
	return u

}
