package main

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"log"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))

func Setup() {
	gothic.Store = store

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	GoogleKey := os.Getenv("GOOGLE_KEY")       // Your Google client ID
	GoogleSecret := os.Getenv("GOOGLE_SECRET") // Your Google client GoogleSecret
	GoogleCallbackURL := "http://localhost:8000/auth/google/callback"

	GithubKey := os.Getenv("GITHUB_ID")        // Your github client ID
	GithubSecret := os.Getenv("GITHUB_SECRET") // Your github client GoogleSecret
	GithubCallbackURL := "http://localhost:8000/auth/github/callback"

	//gothic.Store = sessions.NewCookieStore([]byte("GoogleSecret"))
	gothic.Store = store
	goth.UseProviders(
		google.New(GoogleKey, GoogleSecret, GoogleCallbackURL),
		github.New(GithubKey, GithubSecret, GithubCallbackURL),
	)
}
