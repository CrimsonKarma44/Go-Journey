package main

import (
	"os"
	"time"
)

type Record struct {
	Spent float64
	Alias string
	Day   time.Time
}

func (r *Record) Save() error {
	filename := r.Day + `.txt`
	return os.WriteFile(
		filename,
		[]byte(string(r.Alias)+"\n"+string(r.Spent)),
		0600,
	)
}

//func LoadPage(title string) (*Page, error) {
//	filename := title + ".txt"
//	about, err := os.ReadFile(filename)
//	if err != nil {
//		return nil, err
//	}
//	return &Page{title, about}, err
//}
