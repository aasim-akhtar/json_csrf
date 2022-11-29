package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aasim-akhtar/json_csrf_lab/app"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("main.go")
	startServer()
}

func startServer() {
	r := mux.NewRouter()

	r.HandleFunc("/api/create/lab_1", app.LAB_1).Methods("POST")
	r.HandleFunc("/api/create/lab_2", app.LAB_2).Methods("POST")
	r.HandleFunc("/api/create/lab_3", app.LAB_3).Methods("POST")
	r.HandleFunc("/api/create/lab_4", app.LAB_4).Methods("POST")

	fmt.Println("SERVER STARTED AT PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
