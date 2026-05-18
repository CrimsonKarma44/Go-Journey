package models

import (
	"time"
)

//type FullName struct {
//	Lastname   string `json:"lastname"`
//	Middlename string `json:"middlename"`
//	Firstname  string `json:"firstname"`
//}

type Student struct {
	Name string `json:name`
	//Name        FullName  `json:"name"`
	Age         int       `json:age`
	School      string    `json:school`
	DateOfBirth time.Time `json:date_of_birth`
}

//func (p *Student) save() error {
//	//filename := p.Name + ".txt"
//	return os.WriteFile(filename, p.Body, 0600)
//}
