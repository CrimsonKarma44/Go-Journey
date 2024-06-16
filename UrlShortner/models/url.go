package models

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	IS    int
	Url   string
	Alias string

	created time.Time
	Updated time.Time
}

func (u *Url) Created() time.Time {
	return u.created
}
func (u *Url) Update(newAlias string) {
	u.Alias = newAlias
	u.Updated = time.Now()
}
func (u *Url) Save() *gorm.DB {
}
