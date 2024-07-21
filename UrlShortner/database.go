package main

import (
	"UrlShortner/models"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DNS() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}

// DB is the database connection pool.
//var DB *sql.DB

// SqlInitDB initializes the database connection.
func SqlInitDB() *sql.DB {
	db, err := sql.Open("mysql", DNS())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
	fmt.Println("Database connection established")
	return db
}

func GormInit(model models.Url) *gorm.DB {
	db, err := gorm.Open(mysql.Open(DNS()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	err = db.AutoMigrate(&model)
	return db
}
