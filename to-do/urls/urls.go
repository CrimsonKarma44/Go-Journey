package urls

import (
	"fmt"
	"log"
	"net/http"
	"to-do/handlers"
)

func UrlHandler() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home/", handlers.DisplayHandler)
	//http.HandleFunc("/create", handlers.CreateHandler)

	fmt.Println("Listening on port localhost:8080/home")
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
