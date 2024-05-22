package main

import (
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "karma"
	password = "pass"
	hostname = "127.0.0.1:3306"
	dbname   = "go_server"
)

type Player struct {
	gorm.Model
	username string `gorm:"unique"`
	email    string `gorm:"unique"`
	password string
}

func Dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func InitialMigration(dbName string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(Dsn(dbName)))
	err := db.AutoMigrate(&Player{})
	if err != nil {
		panic(err)
	}

	return db
}
