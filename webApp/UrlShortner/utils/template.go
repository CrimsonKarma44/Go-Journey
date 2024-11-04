package utils

import (
	"UrlShortner/models"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, u []models.Url) {
	t, _ := template.ParseFiles(tmpl + ".html")
	_ = t.Execute(w, u)
}
