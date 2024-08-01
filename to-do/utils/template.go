package utils

import (
	"html/template"
	"net/http"
	"to-do/models"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, u []models.List) {
	t, _ := template.ParseFiles(tmpl + ".html")
	err := t.Execute(w, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
