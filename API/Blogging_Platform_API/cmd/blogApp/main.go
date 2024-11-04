package main

import (
	"Blogging-Platform-API/handler"
	"fmt"
	"log"
	"net/http"
)

const homePage = "http://localhost:8000/posts"

func url() {
	fmt.Println("Starting Server...")
	http.HandleFunc("/posts", handler.CreateViewAllHandler)
	http.HandleFunc("/posts/", handler.UpdateDeleteGetPostHandler)
	fmt.Printf("Server started at %v ...\n", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	url()
}
