package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

// Template Caching
var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

// Validator
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)(/*)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

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
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//this finds every thing after the '/view/'
	//title := r.URL.Path[len("/view/"):]
	//title, err := getTitle(w, r)
	//if err != nil {
	//	return
	//}

	//this loads the .txt page
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, titles string) {
	//titles := r.URL.Path[len("/edit/"):]
	//titles, err := getTitle(w, r)
	//log.Println(titles)
	//if err != nil {
	//	return
	//}

	p, err := loadPage(titles)
	if err != nil {
		p = &Page{Title: titles}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/save/"):]
	//title, err := getTitle(w, r)
	//if err != nil {
	//	return
	//}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func urlValidator(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", urlValidator(viewHandler))
	http.HandleFunc("/edit/", urlValidator(editHandler))
	http.HandleFunc("/save/", urlValidator(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
