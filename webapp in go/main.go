package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//Urlhandler()
	router := mux.NewRouter()
	muxhandler(router)
}

func muxhandler(router *mux.Router) {
	r := mux.NewRouter()

	r.HandleFunc("/", handlerfunc).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Listening on port 8080")
	fmt.Println("Starting server...")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my World!")
}

func Urlhandler() {
	http.HandleFunc("/", handlerfunc)
	fmt.Println("Listening on port 8080")
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
