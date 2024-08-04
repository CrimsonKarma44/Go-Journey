package handlers

import (
	"fmt"
	"github.com/markbates/goth/gothic"
	"github/utils"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println(":5000/home")
	utils.RenderTemplate(w, "home", nil)
}

func GitAuth(w http.ResponseWriter, r *http.Request) {
	log.Println(":5000/auth/github")
	gothic.BeginAuthHandler(w, r)
}

func GitAuthCallback(w http.ResponseWriter, r *http.Request) {
	log.Println(":5000/auth/github/callback")
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%#v", user)
}
