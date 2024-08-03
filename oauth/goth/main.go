//package main
//
//import (
//	"fmt"
//	"github.com/joho/godotenv"
//	"github.com/markbates/goth"
//	"github.com/markbates/goth/providers/google"
//	"log"
//	"net/http"
//	"os"
//)
//
//func urls() {
//	http.HandleFunc("/", Home)
//	http.HandleFunc("/auth/{provider}/callback", CallBackHandler)
//	http.HandleFunc("/auth/{provider}", AuthHandler)
//
//	fmt.Println("Listening on port localhost:8080/home")
//	fmt.Println("Starting server...")
//	log.Fatal(http.ListenAndServe("localhost:8080", nil))
//}
//
//func main() {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	goth.UseProviders(
//		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), CallbackURL, "email", "profile"),
//	)
//
//	urls()
//}

//package main
//
//import (
//	"fmt"
//	"github.com/joho/godotenv"
//	"log"
//	"net/http"
//	"os"
//
//	"github.com/gorilla/mux"
//	"github.com/markbates/goth"
//	"github.com/markbates/goth/gothic"
//	"github.com/markbates/goth/providers/google"
//)
//
//func main() {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//	key := os.Getenv("GOOGLE_KEY")       // Your Google client ID
//	secret := os.Getenv("GOOGLE_SECRET") // Your Google client secret
//	callbackURL := "http://localhost:8080/auth/google/callback"
//
//	goth.UseProviders(
//		google.New(key, secret, callbackURL, "email", "profile"),
//	)
//
//	r := mux.NewRouter()
//
//	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "<a href='/auth/google'>Log in with Google</a>")
//	})
//
//	r.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
//		user, err := gothic.CompleteUserAuth(w, r)
//		if err != nil {
//			fmt.Fprintln(w, err)
//			return
//		}
//		fmt.Fprintf(w, "User: %#v", user)
//	})
//
//	r.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
//		gothic.BeginAuthHandler(w, r)
//	})
//
//	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
//		gothic.Logout(w, r)
//		fmt.Fprintln(w, "Logged out!")
//	})
//
//	http.ListenAndServe(":8080", r)
//}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:3000/auth/google/callback"),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize itself in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	//openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), "http://localhost:3000/auth/openid-connect/callback", os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
	//if openidConnect != nil {
	//	goth.UseProviders(openidConnect)
	//}

	m := map[string]string{
		"google": "Google",
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, user)
	})

	p.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(userTemplate)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, providerIndex)
	})

	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
