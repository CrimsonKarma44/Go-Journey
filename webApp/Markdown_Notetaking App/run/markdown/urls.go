package main

import (
	"log"
	"net/http"
)

func url() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
