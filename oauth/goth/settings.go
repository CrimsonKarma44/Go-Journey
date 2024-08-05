package main

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"log"
	"os"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func Setup() {
	gothic.Store = store

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("GOOGLE_KEY")       // Your Google client ID
	secret := os.Getenv("GOOGLE_SECRET") // Your Google client secret
	callbackURL := "http://localhost:8000/auth/google/callback"
	gothic.Store = sessions.NewCookieStore([]byte("secret"))
	goth.UseProviders(
		google.New(key, secret, callbackURL),
	)
}
