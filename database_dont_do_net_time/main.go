package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "karma"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "go_server"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

// Hashes a password using bcrypt
func hashPassword(password string) (string, error) {
	// Generate a salt with a cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Compares a hashed password with a plaintext password
func comparePasswords(hashedPassword, enteredPassword string) error {
	// Compare the hashed password with the plaintext password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
}

func main() {
	db, _ := gorm.Open(mysql.Open(dsn(dbname)))
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("An error occured at migration: %s", err)
	}
	//db.Create(&User{Username: "James", Password: "123456"})
	//db.Model(&User{}).Where("Username = ?", "James").Update("Password", hashed_password)

	var username string
	var password string
	fmt.Println("Create Username:")
	fmt.Scan(&username)
	fmt.Println("Create Password:")
	fmt.Scan(&password)

	hashedPassword, _ := hashPassword(password)

	db.Create(&User{Username: username, Password: hashedPassword})

	fmt.Println("User created successfully")
	var condition string
	fmt.Println("Login? yes or no")
	fmt.Scan(&condition)

	if condition == "yes" {
		for {
			fmt.Println("Password:")
			fmt.Scan(&password)
			var user User

			db, err := sql.Open("mysql", dsn(dbname))
			if err != nil {
				panic(err)
			}
			defer db.Close()
			query, err := db.Query("SELECT password FROM users WHERE username = ?", username)
			if err != nil {
				panic(err)
			}
			//db.Where("Username = ?", username).First(&user)
			fmt.Println(password)
			query.Next()
			err = query.Scan(&hashedPassword)
			if err != nil {
				panic(err)
			}
			fmt.Println(hashedPassword)
			err = comparePasswords(hashedPassword, password)
			if err != nil {
				fmt.Printf("Wrong password: '%s' not the same as '%s'", password, user.Password)
			} else {
				fmt.Printf("User %s successfully Logged-in!", user.Username)
				return
			}
		}
	} else {
		fmt.Println("Bye")
		return
	}

	//db.Update()
	//db.Delete(&User{}, 1)

	//db, err := sql.Open("mysql", dsn(dbname))
	//if err != nil {
	//	log.Printf("Error %s when opening DB\n", err)
	//	return
	//} else {
	//	fmt.Println("We are connected to the DB")
	//}
	//defer db.Close()
	//
	//_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	//if err != nil {
	//	log.Printf("Error %s when creating DB\n", err)
	//}
	//fmt.Printf("Created DB %s\n", dbname)

}

//func main() {
//	fmt.Println("Go MySQL Tutorial")
//
//	// Open up our database_dont_do_net_time connection.
//	// I've set up a database_dont_do_net_time on my local machine using phpmyadmin.
//	// The database_dont_do_net_time is called testDb
//	db, err := sql.Open("mysql", "karma:pass@tcp(127.0.0.1:3306)/eventify")
//
//	// if there is an error opening the connection, handle it
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// defer the close till after the main function has finished
//	// executing
//	defer db.Close()
//}
