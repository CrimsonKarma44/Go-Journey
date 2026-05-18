package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"verhiser/models"
)

func GormPostgresqlInit() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DNSPostgresql()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		return nil, err
	}

	return db, err
}

func DNSPostgresql() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}
