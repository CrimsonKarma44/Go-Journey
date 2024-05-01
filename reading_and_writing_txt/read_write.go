package main

import (
	"fmt"
	"os"
)

type Page struct {
	title string
	about []byte
}

func (p *Page) save() error {

	filename := p.title + ".txt"
	return os.WriteFile(filename, p.about, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	about, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, about}, err
}

func main() {
	p1 := &Page{title: "TestPage", about: []byte("This is a simple Page a very simple one .")}
	error := p1.save()
	if error != nil {
		fmt.Println(error)
	} else {
		p2, _ := loadPage("TestPage")
		fmt.Println(string(p2.about))
	}
	fmt.Println(len("/view/"))
}
