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
	rootMux.HandleFunc("/test", handlers.Test)
	rootMux.Handle("/auth/", authUrl)
	//http.HandleFunc("/home", handlers.Home)
	//http.HandleFunc("/test", handlers.Test)
	//
	//http.HandleFunc("/auth/github", handlers.GitAuth)
	//http.HandleFunc("/auth/github/callback", handlers.GitAuthCallback)

	log.Println("Starting server on port http://localhost:8000/home")
	log.Println("Listening on port 8000")

	log.Fatal(http.ListenAndServe(":8000", rootMux))
	//log.Fatal(http.ListenAndServe(":8000", nil))
}
