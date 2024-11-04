package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	username = "karma"
	password = "pass"
	hostname = "127.0.0.1:3306"
	dbname   = "eventify"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

type Censor struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Region    string `gorm:"not null"`
	Age       int    `gorm:"not null"`
}

func InitialMigration(dbName string) *gorm.DB {
	db, _ := gorm.Open(mysql.Open(dsn(dbName)))
	err := db.AutoMigrate(&Censor{})
	if err != nil {
		panic(err)
	}

	return db
}

func AddEntry(db *gorm.DB) {
	var censor Censor

	fmt.Println("First Name: ")
	fmt.Scan(&censor.FirstName)
	fmt.Println("Last Name: ")
	fmt.Scan(&censor.LastName)
	fmt.Println("Region: ")
	fmt.Scan(&censor.Region)
	fmt.Println("Age: ")
	fmt.Scan(&censor.Age)

	db.Create(&censor)
	fmt.Println("Entry created!")
}

func FetchCensors(dbName string) []Censor {
	db, _ := gorm.Open(mysql.Open(dsn(dbName)))
	var censors []Censor
	db.Find(&censors)

	for _, censor := range censors {
		fmt.Println("Firstname:`", censor.FirstName, "`")
		fmt.Println("LastName:", censor.LastName)
		fmt.Println("Region:", censor.Region)
		fmt.Println("Age:", censor.Age)
		fmt.Println()
	}
	return censors
}

func main() {
	//InitialMigration(dbname)
	//AddEntry(db)
	//FetchCensors(dbname)

	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query, err := db.Query("SELECT first_name, last_name, region, age FROM censors WHERE deleted_at IS NULL")
	if err != nil {
		panic(err)
	}

	for query.Next() {
		var censor Censor
		err := query.Scan(&censor.FirstName, &censor.LastName, &censor.Region, &censor.Age)
		if err != nil {
			panic(err.Error())
		}
		log.Println(censor.LastName, censor.FirstName, censor.Age, censor.Region)
	}

	//query2, err := db.
}
