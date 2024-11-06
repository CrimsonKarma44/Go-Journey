package main

import (
	"Blogging_Platform_API/handler"
	"Blogging_Platform_API/utility"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	err := initialize(utility.DNS())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connection established")
	url()
}

func initialize(dns string) error {
	fmt.Println("Initializing Database...")
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	return db.Ping()
}
