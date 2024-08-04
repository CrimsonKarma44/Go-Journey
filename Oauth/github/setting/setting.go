package setting

import (
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"log"
	"os"
)

//var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	GithubID := os.Getenv("GITHUBID")
	GithubSecret := os.Getenv("GITHUBSECRET")
	CallbackURLGithub := "http://localhost:5000/auth/github/callback"

	//gothic.Store = store
	goth.UseProviders(
		github.New(GithubID, GithubSecret, CallbackURLGithub, "user:email"),
	)
}
