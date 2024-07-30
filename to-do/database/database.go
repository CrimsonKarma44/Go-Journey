package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"to-do/models"
)

func DNS() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}
func SQLInit() *sql.DB {
	db, err := sql.Open("mysql", DNS())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection established")

	return db
}
func GormInit() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DNS()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")
	err = db.AutoMigrate(&models.User{}, &models.List{})
	if err != nil {
		log.Fatalf("Error auto migrate lists: %v", err)
		return nil, err
	}

	log.Println("Migration complete")
	return db, nil
}
