package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model

	EventId           []Event `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
	Name              string  `gorm:"varchar(255)"`
	ticketType        string  `gorm:"varchar(255)"`
	price             float32 `gorm:"decimal(10,2)"`
	quantityAvailable int     `gorm:"int(11)"`
	availStatus       bool
}

func (t *Ticket) __str__() string {
	return fmt.Sprintf("Ticket %s{event_name} - %s{self.name}", t.EventId)
}
