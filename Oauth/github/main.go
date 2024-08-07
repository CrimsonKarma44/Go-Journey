package main

import (
	"github/setting"
	"github/urls"
)

//var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	setting.Init()
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}
	//GithubID := os.Getenv("GITHUB_ID")
	//GithubSecret := os.Getenv("GITHUB_SECRET")
	//CallbackURLGithub := "http://localhost:5000/auth/github/callback"
	//
	//gothic.Store = store
	//goth.UseProviders(
	//	github.New(GithubID, GithubSecret, CallbackURLGithub, "user:email"),
	//)
	urls.UrlHandler()
}
