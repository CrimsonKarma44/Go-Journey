package models

import "gorm.io/gorm"

type EventPlanner struct {
	gorm.Model
}

type EventRegular struct {
	gorm.Model
}
