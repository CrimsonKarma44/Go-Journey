package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
	"html/template"
	"log"
	"net/http"
)

func urlHandler() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//RenderTemplate(w, "home", nil)
		fmt.Fprintln(w, "<a href='/auth/google'>Log in with Google</a>")
	})

	r.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		//gothic.StoreInSession("userEmail", user.Email, r, w)
		//gothic.StoreInSession("userIDToken", user.IDToken, r, w)
		//gothic.StoreInSession("userAccessToken", user.AccessToken, r, w)
		//gothic.StoreInSession("userAccessTokenSecret", user.AccessTokenSecret, r, w)

		sessions, _ := store.Get(r, "TestUser")
		sessions.Values["user"] = user
		sessions.Save(r, w)

		//fmt.Fprintf(w, "%#v", user)
		http.Redirect(w, r, "/test", http.StatusFound)
	})

	r.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[len("/auth/"):]
		log.Println("/auth/" + title)

		gothic.BeginAuthHandler(w, r)
	})

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "TestUser")
		user, _ := session.Values["user"]
		if user == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			session.Options.MaxAge = -1 // Invalidate the session
			session.Save(r, w)
			gothic.Logout(w, r)

			log.Println("Logged out!")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})

	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		//user, _ := gothic.CompleteUserAuth(w, r)
		//type userInfo struct {
		//	Email                 string
		//	UserIDToken           string
		//	UserAccessToken       string
		//	UserAccessTokenSecret string
		//}
		//user := userInfo{}
		//user.Email, _ = gothic.GetFromSession("userEmail", r)
		//user.UserIDToken, _ = gothic.GetFromSession("userIDToken", r)
		//user.UserAccessToken, _ = gothic.GetFromSession("userAccessToken", r)
		//user.UserAccessTokenSecret, _ = gothic.GetFromSession("userAccessTokenSecret", r)

		sessions, _ := store.Get(r, "TestUser")
		user, _ := sessions.Values["user"]
		if user == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		}

		//fmt.Fprintf(w, "%#v", user)

		RenderTemplate(w, "templates/home", user)
	})

	fmt.Println("Starting Server at port \"localhost:8000\" .....")

	http.ListenAndServe(":8000", r)

}

func main() {
	Setup()
	urlHandler()
}

func RenderTemplate(w http.ResponseWriter, tmpl string, u interface{}) {
	t, _ := template.ParseFiles(tmpl + ".html")
	_ = t.Execute(w, u)
}
