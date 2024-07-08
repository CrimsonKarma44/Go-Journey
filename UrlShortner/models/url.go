package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	IS    int
	Url   string
	Alias string
}

func (u *Url) Created() string {
	return string(u.CreatedAt.Day()) + string(u.CreatedAt.Hour()) + ":" + string(u.CreatedAt.Minute()) + ":" + string(u.CreatedAt.Second())
}

//	func (u *Url) Update(newAlias string) {
//		u.Alias = newAlias
//		u.Updated = time.Now()
//	}
//func (u *Url) Save() *gorm.DB {
//}

var db *gorm.DB

func RevealUrl() []Url {
	var urls []Url
	db.Find(&urls)
	return urls
}
