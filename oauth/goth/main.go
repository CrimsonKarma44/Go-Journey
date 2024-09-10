package main

import (
	"html/template"
	"net/http"
)

func main() {
	Setup()
	UrlHandler()
}

func RenderTemplate(w http.ResponseWriter, tmpl string, u interface{}) {
	t, _ := template.ParseFiles(tmpl + ".html")
	_ = t.Execute(w, u)
}
