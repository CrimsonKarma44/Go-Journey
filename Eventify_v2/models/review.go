package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	EventId []Event `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
	Email   string  `gorm:"type:varchar(200)"`
	Content string  `gorm:"type:text"`
	Rate    string  `gorm:"type:varchar(200)"`
}

func (r *Review) __str__() string {
	return fmt.Sprintf("%s", r.EventId)
}
