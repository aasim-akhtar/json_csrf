package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	UserName string
	Role     string
}

func (u User) createUser(w http.ResponseWriter) {
	fmt.Fprintf(w, "\nUser Created.\n%s is now a(n) %s", u.UserName, u.Role)
}

// JSON CSRF with no Padding & Content-Type Validation
func LAB_1(w http.ResponseWriter, r *http.Request) {
	var p User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "LAB 1 Solved")
	p.createUser(w)
}

// JSON CSRF with JSON Request Body Validation (No Unknown Fields Allowed)
func LAB_2(w http.ResponseWriter, r *http.Request) {
	var p User

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		fmt.Fprintf(w, "Error %+v\nFailed to create user: %+v ", err, p.UserName)
		return
	}
	p.createUser(w)
}

func LAB_3(w http.ResponseWriter, r *http.Request) {}
func LAB_4(w http.ResponseWriter, r *http.Request) {}
