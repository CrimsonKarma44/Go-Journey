package main

import (
	"fmt"
	"log"
	"net/http"
)

func url() {
	log.Fatal(http.ListenAndServe(":3000", http.FileServer(http.Dir("assets/"))))
}

func main() {
	fmt.Println("love")
}
