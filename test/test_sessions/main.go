package main

import (
	"fmt"
	"net/http"
	"time"
)

func urlHandler() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/set-cookie", setCookieHandler)
	http.HandleFunc("/read-cookie", readCookieHandler)
	http.HandleFunc("/delete-cookie", deleteCookieHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	urlHandler()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<a href="/set-cookie">Set Cookie</a><br>
                     <a href="/read-cookie">Read Cookie</a><br>
                     <a href="/delete-cookie">Delete Cookie</a>`)
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour) // Cookie expiration time
	cookie := http.Cookie{
		Name:     "myCookie",
		Value:    "Hello, World!",
		Expires:  expiration,
		HttpOnly: true, // Set HttpOnly to true for security
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie set!")
}

func readCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("myCookie")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No cookie found!")
		} else {
			fmt.Fprintln(w, "Error reading cookie:", err)
		}
		return
	}
	fmt.Fprintf(w, "Cookie value: %s\n", cookie.Value)
}

func deleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "myCookie",
		Value:    "",
		Expires:  time.Unix(0, 0), // Set expiration date in the past to delete the cookie
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie deleted!")
}
