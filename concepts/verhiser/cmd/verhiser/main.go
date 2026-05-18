package main

import (
	"log"
	"net/http"
	"verhiser/database"
)

func main() {
	database.GormPostgresqlInit()
	url()
}

func url() {
	mux := http.NewServeMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
