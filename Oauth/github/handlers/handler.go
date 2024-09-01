package handlers

import (
	"context"
	"fmt"
	"github.com/markbates/goth/gothic"
	"github/utils"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println(":8000/home")
	fmt.Fprintf(w, "<a href='/auth/github'>login</a> to github")
	fmt.Fprintf(w, "<a href='/test'>test</a>")
	//utils.RenderTemplate(w, "home", nil)
}
func Test(w http.ResponseWriter, r *http.Request) {
	log.Println(":8000/home")
	utils.RenderTemplate(w, "home", nil)
}

func GitAuth(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Path[len("/auth/"):]
	log.Println(":8000/auth/", provider)
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothic.BeginAuthHandler(w, r)
}

func GitAuthCallback(w http.ResponseWriter, r *http.Request) {
	log.Println(":8000/auth/github/callback")
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%#v", user)
}
