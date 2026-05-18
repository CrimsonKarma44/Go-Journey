package handler

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserHandler struct {
	DB *gorm.DB
}

func (u *UserHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		user := database.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
			Role:     r.FormValue("role"),
		}
		err := user.Create(u.DB)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		err := user.RetrieveEmail(u.DB, email)
		if err != nil || utils.ComparePasswords(user.Password, password) != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token, err := utils.GenerateToken(email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Authorization", token)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (u *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Authorization", "")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User Logged Out"))
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (u *UserHandler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	if r.Method == http.MethodGet {
		productID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user.RetrieveID(u.DB, uint(productID))
		w.WriteHeader(http.StatusOK)
		jsonText, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(jsonText)
		return
	}
}
