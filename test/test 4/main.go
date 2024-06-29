package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(20);not null"`
	SchoolID uint   // Foreign key field
	School   School // Association with School
}

// School represents a school in the system
type School struct {
	ID         uint   `gorm:"primaryKey"`
	SchoolName string `gorm:"type:varchar(20);not null"`
	Students   []User `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE"`
}

func InitDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		"karma",
		"password",
		"localhost",
		"3306",
		"go_server",
	)
	db, _ := gorm.Open(mysql.Open(dsn))
	err := db.AutoMigrate(&User{}, &School{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func url() {
	http.HandleFunc("/{name}/{surname}/", func(writer http.ResponseWriter, request *http.Request) {
		name := request.PathValue("name")
		surname := request.PathValue("surname")
		log.Println("/" + name + " " + surname)

		_, err := writer.Write([]byte(name + " " + surname))
		if err != nil {
			return
		}
	})
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	db := InitDb()
	populateDatabase(db)
	printData(db)
	url()
}

func populateDatabase(db *gorm.DB) {
	// Create some schools
	schools := []School{
		{SchoolName: "Greenwood High"},
		{SchoolName: "Riverside Academy"},
	}
	db.Create(&schools)

	// Create some users and associate them with the schools
	users := []User{
		{Name: "Alice", Email: "alice@example.com", SchoolID: schools[0].ID},
		{Name: "Bob", Email: "bob@example.com", SchoolID: schools[0].ID},
		{Name: "Charlie", Email: "charlie@example.com", SchoolID: schools[1].ID},
	}
	db.Create(&users)
}

func printData(db *gorm.DB) {
	var schools []School
	db.Preload("Students").Find(&schools)

	for _, school := range schools {
		fmt.Printf("School: %s\n", school.SchoolName)
		for _, student := range school.Students {
			fmt.Printf("  Student: %s, Email: %s\n", student.Name, student.Email)
		}
	}
}
