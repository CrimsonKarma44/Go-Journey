package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func UrlHandler() {
	r := mux.NewRouter()

	r.HandleFunc("/", Home)

	r.HandleFunc("/auth/{provider}/callback", CallBackHandler)
	r.HandleFunc("/auth/{provider}", AuthHandler)

	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/test", TestHandler)

	fmt.Println("Starting Server at port \"http://localhost:8000\" .....")
	http.ListenAndServe(":8000", r)

}
