package main

import (
	"Personal_Budget_Tracker/handler"
	"Personal_Budget_Tracker/models"
	"fmt"
)

func main() {
	//var duration time.Duration
	//instance := 2
	//newTime := time.Date(2006, 01, 02, 15, 04, 05, 22, time.UTC)
	//duration = time.Duration(instance*24) * (time.Hour)
	//change := newTime.Add(duration)
	//fmt.Println(newTime)
	//fmt.Println(change)
	//fmt.Println("Duration: ", duration)

	var value string
	var newBudget models.BudgetExpense
	handler.Init()
	for {
		//handler.DisplayBudget(newBudget)
		fmt.Printf(">> ")
		fmt.Scanln(&value)
		switch value {
		case "create":
			newBudget = handler.CreateBudget()
			handler.DisplayBudget(newBudget)
			continue
		case "display":
			handler.DisplayBudget(newBudget)
			continue
		default:
			continue
		}
	}
}
