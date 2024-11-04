package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB is the database connection pool.
var DB *sql.DB

func DNS() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}

// InitDB initializes the database connection.
func InitDB() {
	DB, err := sql.Open("mysql", DNS())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	fmt.Println("Database connection established")
}
