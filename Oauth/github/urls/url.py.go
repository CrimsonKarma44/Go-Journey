package urls

import (
	"github/handlers"
	"log"
	"net/http"
)

func UrlHandler() {
	authUrl := http.NewServeMux()
	authUrl.HandleFunc("/auth/github", handlers.GitAuth)
	authUrl.HandleFunc("/auth/github/callback", handlers.GitAuthCallback)

	rootMux := http.NewServeMux()
	rootMux.HandleFunc("/home", handlers.Home)
	rootMux.Handle("/auth/", authUrl)

	log.Println("Starting server on port http://localhost:5000/home")
	log.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", rootMux))
}
