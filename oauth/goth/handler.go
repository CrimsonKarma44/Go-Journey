package main

import (
	"fmt"
	"github.com/markbates/goth/gothic"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<a href='/auth/google'>Log in with Google</a></br>")
	fmt.Fprintln(w, "<a href='/auth/github'>Log in with Github</a>")

	log.Println("/")
}

func CallBackHandler(w http.ResponseWriter, r *http.Request) {
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

	title := r.URL.Path[len("/auth/"):]
	log.Println("/auth/" + title + " authorized...")
	//fmt.Fprintf(w, "%#v", user)
	http.Redirect(w, r, "/test", http.StatusFound)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	//gothic.BeginAuthHandler(w, r)
	title := r.URL.Path[len("/auth/"):]
	log.Println("/auth/" + title + " authorizing...")

	gothic.BeginAuthHandler(w, r)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
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
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Println("/test")
	RenderTemplate(w, "templates/home", user)
}
