package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	TicketId      []Ticket `gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE"`
	Price         float32
	Email         string `gorm:"type:varchar(255)"`
	Code          string `gorm:"type:varchar(255)"`
	PhoneNumber   string `gorm:"type:varchar(20)"`
	TransactionId int64
	Present       bool
}

func (p *Payment) __str__() string {
	return fmt.Sprintf(
		"%s %s payment for %s",
		p.Email,
		p.Price,
		p.TicketId,
	)
}
