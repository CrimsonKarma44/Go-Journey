package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	InitDB()

	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.Handle("/welcome", AuthMiddleware(http.HandlerFunc(WelcomeHandler))).Methods("GET")

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
