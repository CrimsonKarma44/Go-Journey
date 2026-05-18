package model

import (
	"os"

	"github.com/joho/godotenv"
)

type Data_values struct {
	User               string
	Password           string
	Host               string
	Db_Name            string
	Ports              string
	SECRET_KEY_Access  []byte
	SECRET_KEY_Refresh []byte
}

func (d Data_values) Load() Data_values {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return Data_values{
		User:               os.Getenv("DB_USER"),
		Password:           os.Getenv("DB_PASSWORD"),
		Host:               os.Getenv("DB_HOST"),
		Db_Name:            os.Getenv("DB_NAME"),
		Ports:              os.Getenv("DB_PORT"),
		SECRET_KEY_Access:  []byte(os.Getenv("JWT_SECRET_Key_Access")),
		SECRET_KEY_Refresh: []byte(os.Getenv("JWT_SECRET_Key_Refresh")),
	}
}
