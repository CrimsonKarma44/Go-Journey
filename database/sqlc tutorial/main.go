package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	sqlc "sqlc/database"

	_ "github.com/go-sql-driver/mysql"
)

func DNS() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}
func run() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", DNS())
	if err != nil {
		return err
	}

	queries := sqlc.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateAuthor(ctx, sqlc.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedAuthorID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorID)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorID)
	if err != nil {
		return err
	}
	//fmt.Println(fetchedAuthor)

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
