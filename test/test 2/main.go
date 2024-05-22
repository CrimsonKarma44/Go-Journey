package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "karma"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "go_server"
)

type Player struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

func (p *Player) Save(db *gorm.DB) (*gorm.DB, error) {
	p.Password, _ = hashPassword(p.Password)

	err := db.Create(&p)
	if err != nil {
		//panic(err)
		return nil, err.Error
	}
	return db, nil
}

func Dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func InitialMigration(dbName string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(Dsn(dbName)))
	err = db.AutoMigrate(&Player{})
	if err != nil {
		fmt.Println("A Problem occured during initial migration")
		panic(err)
	}

	return db
}

//		response = Player{Username: "Karma", Email: "karma@gmail.com", Password: "pass"}
//		jsonResponse, err := json.Marshal(response)
//		if err != nil {
//			http.Error(w, "Error converting response to JSON", http.StatusInternalServerError)
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(jsonResponse)
//	} else {
func registerPlayer(w http.ResponseWriter, r *http.Request) {
	var response Player
	db, err := gorm.Open(mysql.Open(Dsn(dbname)))
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return
		}
		_, err = response.Save(db)
		if err != nil {
			http.Error(w, "Unable to register player", http.StatusInternalServerError)
		} else {
			w.Write([]byte("success"))
		}
	}
}

func urlhandler() {
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/register/", registerPlayer)
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func main() {
	// initializing the database
	_ = InitialMigration(dbname)
	urlhandler()

	//starting the server

}

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
