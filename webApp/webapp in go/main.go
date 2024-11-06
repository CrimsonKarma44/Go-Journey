package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//Urlhandler()
	muxhandler()
}

func muxhandler() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlerfunc).Methods("GET")
	//http.Handle("/", r)

	fmt.Println("Listening on port 8080")
	fmt.Println("Starting server...")

	//log.Fatal(http.ListenAndServe("localhost:8080", nil))
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	log.Println(":8080/")
	fmt.Fprintf(w, "Hello my World!")
}

func Urlhandler() {
	http.HandleFunc("/", handlerfunc)
	fmt.Println("Listening on port 8080")
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
