package setting

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"log"
	"os"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	GithubID := os.Getenv("GITHUB_ID")
	GithubSecret := os.Getenv("GITHUB_SECRET")
	CallbackURLGithub := "http://localhost:8000/auth/github/callback"

	gothic.Store = store
	goth.UseProviders(
		github.New(GithubID, GithubSecret, CallbackURLGithub),
	)
}
