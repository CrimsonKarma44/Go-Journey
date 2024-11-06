package main

import (
	"Unit-Converter/handlers"
	"fmt"
	"log"
	"net/http"
)

func urlHandler() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", handlers.HomeHandler)
	fmt.Println("...Starting at http://localhost:8000/home")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("... Starting Unit Converter ...")
	fmt.Println()
	urlHandler()
}
