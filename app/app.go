package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"

	"github.com/go-openapi/runtime/middleware/header"
)

const (
	L1 = "Lab 1"
	L2 = "Lab 2"
	L3 = "Lab 3"
	L4 = "Lab 4"
)

type User struct {
	UserName string `json:"UserName" validate: "alphanum,excludesall=0x3d"`
	Role     string `json:"Role" validate: "alphanum,excludesall=0x3d"`
}

func (u User) createUser(w http.ResponseWriter, l string) {
	fmt.Fprint(w, l, " Solved")
	fmt.Fprintf(w, "\nUser Created.\n%s is now a(n) %s", u.UserName, u.Role)
	fmt.Printf("\n%s Solved\nUser Created.\n%s is now a(n) %s", l, u.UserName, u.Role)
}

// JSON CSRF without any Validation
func LAB_1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nLab 1 Initiated")
	var p User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.createUser(w, L1)
}

// JSON CSRF with JSON Request Body Validation (No Unknown Fields Allowed)
func LAB_2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nLab 2 Initiated")
	var p User

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		fmt.Fprintf(w, "Error %+v\nFailed to create user: %+v ", err, p.UserName)
		return
	}
	p.createUser(w, L2)
}

// JSON CSRF with Strict POST Body Validation & no Content-Type Validation
func LAB_3(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\nLab 3 Initiated")
	var p User

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		fmt.Fprintf(w, "Error %+v\nFailed to create user: %+v ", err, p.UserName)
		return
	}

	values := reflect.ValueOf(p)
	for i := 0; i < values.NumField(); i++ {
		m, _ := regexp.MatchString("^[a-zA-Z]+$", values.Field(i).String())
		if m != true {
			fmt.Fprintf(w, "Invalid Special Characters or Numbers in Body \nFailed to create user")
			return
		}

	}

	// validate := validator.New()
	// err = validate.Struct(p)
	// // validationErrors := err.(validator.ValidationErrors)
	// if err != nil {

	// 	// this check is only needed when your code could produce
	// 	// an invalid value for validation such as interface with nil
	// 	// value most including myself do not usually have code like this.
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {

	// 		fmt.Println(err.Namespace())
	// 		fmt.Println(err.Field())
	// 		fmt.Println(err.StructNamespace())
	// 		fmt.Println(err.StructField())
	// 		fmt.Println(err.Tag())
	// 		fmt.Println(err.ActualTag())
	// 		fmt.Println(err.Kind())
	// 		fmt.Println(err.Type())
	// 		fmt.Println(err.Value())
	// 		fmt.Println(err.Param())
	// 		fmt.Println(err)
	// 	}

	// 	// from here you can create your own error messages in whatever language you wish
	// 	return
	// }

	p.createUser(w, L3)
}

// JSON CSRF with Content-Type Header Validation
func LAB_4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nLab 4 Initiated")
	var p User
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.createUser(w, L4)
}
