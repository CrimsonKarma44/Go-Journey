package main

import (
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"log"
	"os"
)

var CallbackURL = "http://localhost:8080/auth/google/callback"

func Setup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), CallbackURL, "email", "profile"),
	)

}
