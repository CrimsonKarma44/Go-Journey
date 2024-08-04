package utils

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	t, err := template.ParseFiles("templates/" + name + ".html")
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}
