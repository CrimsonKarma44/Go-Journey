package helper

import (
	"Unit-Converter/models"
	"html/template"
	"net/http"
)

// Template Caching
var templates = template.Must(template.ParseFiles("templates/home.html", "templates/result.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, f *models.Form) {
	err := templates.ExecuteTemplate(w, tmpl+".html", f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
