package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aasim-akhtar/json_csrf_lab/app"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("main.go")
	startServer()
}

func startServer() {
	r := mux.NewRouter()
	// @todo ACAC with ACAO set to * will not work
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET"})
	// Discuss why 3600 is not reflected in ACMA
	ttl := handlers.MaxAge(3600)

	//  @todo ACAO is prohibited from using a star * for requests with credentials.
	origins := handlers.AllowedOrigins([]string{"null", "http://burpsuite"})
	// @todo @issue ACAH is not responding with Accept.
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type"})

	r.HandleFunc("/api/create/lab_1", app.LAB_1).Methods("POST")
	r.HandleFunc("/api/create/lab_2", app.LAB_2).Methods("POST")
	r.HandleFunc("/api/create/lab_3", app.LAB_3).Methods("POST")
	r.HandleFunc("/api/create/lab_4", app.LAB_4).Methods("POST")

	fmt.Println("SERVER STARTED AT PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(credentials, methods, origins, headers, ttl)(r)))
}
