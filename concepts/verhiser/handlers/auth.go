package handlers

import (
	"gorm.io/gorm"
	"net/http"
)

type Auth struct {
	DB *gorm.DB
}

func (a Auth) Login(w http.ResponseWriter, r *http.Request) {

}
func (a Auth) Signup(w http.ResponseWriter, r *http.Request) {

}
