package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	_ = add(1, 2)
	var u user
	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}
	if session, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[session.un]
	}
	return u

}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	session := dbSessions[c.Value]
	_, ok := dbUsers[session.un]
	return ok

}
func add(x int, y int) int {
	return x + y
}
