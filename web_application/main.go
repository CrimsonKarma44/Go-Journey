package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

//func create(header string) {
//	p1 := &Page{Title: header, Body: []byte("This is a simple Page a very not so simple one .")}
//	error := p1.save()
//	if error != nil {
//		fmt.Println(error)
//	} else {
//		p2, _ := loadPage("TestPage")
//		fmt.Println(string(p2.Body))
//	}
//}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	about, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, about}, err
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//this finds every thing after the '/view/'
	title := r.URL.Path[len("/view/"):]
	//this loads the .txt page
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	titles := r.URL.Path[len("/edit/"):]
	p, err := loadPage(titles)
	if err != nil {
		p = &Page{Title: titles}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
