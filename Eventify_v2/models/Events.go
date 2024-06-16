package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	User        []EventPlanner `gorm:"foreignKey:EventPlannerID;constraint:OnDelete:CASCADE"`
	Name        string         `gorm:"not null;size:255"`
	Description string         `gorm:"size:500"`
	Type        EventTypes     `gorm:"type:ENUM('Concert', 'Communities', 'Classes', 'Party', 'Sport');default:'Party'"`
	Location    string         `gorm:"not null;size:255"`
	Start       time.Time
	End         time.Time
}

func (event *Event) __str__() string {
	return fmt.Sprintf(event.Name)
}

type EventTypes string

const (
	Concert     EventTypes = "Concert"
	Communities EventTypes = "Communities"
	Classes     EventTypes = "Classes"
	Party       EventTypes = "Party"
	Sport       EventTypes = "Sport"
)
