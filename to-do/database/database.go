package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"to-do/models"
)

func DNSMysql() string {
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

func SQLInit() *sql.DB {
	db, err := sql.Open("mysql", DNSMysql())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection established")

	return db
}
func GormMysqlInit() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DNSMysql()), &gorm.Config{})
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
func GormPostgresqlInit() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DNSPostgresql()), &gorm.Config{})
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
