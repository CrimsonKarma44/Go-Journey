package urls

import (
	"UrlShortner/views"
	"fmt"
	"log"
	"net/http"
)

func UrlHandler() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/home", views.DisplayHandler)
	http.HandleFunc("/home/delete/", views.DeleteHandler)
	http.HandleFunc("/home/copy/", views.CopyHandler)

	fmt.Println("Listening on port http://localhost:8000/home")
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
