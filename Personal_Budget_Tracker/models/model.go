package models

import "time"

type BudgetExpense struct {
	Budget    float64
	Expense   []Money
	Duration  time.Duration
	CreatedAt time.Time
}

type Money struct {
	Title string
	Price float64
}

func (b *BudgetExpense) Save() error {
	return nil
}
