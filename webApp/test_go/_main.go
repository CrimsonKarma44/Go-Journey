package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

var users = make(map[string]string)

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if _, exists := users[user.Username]; exists {
        http.Error(w, "Username already exists", http.StatusBadRequest)
        return
    }

    users[user.Username] = user.Password
    fmt.Fprintf(w, "User %s registered successfully", user.Username)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    password, exists := users[user.Username]
    if !exists || password != user.Password {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    fmt.Fprintf(w, "User %s logged in successfully", user.Username)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "User logged out successfully")
}

func main() {
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/logout", logoutHandler)

    fmt.Println("Server running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
