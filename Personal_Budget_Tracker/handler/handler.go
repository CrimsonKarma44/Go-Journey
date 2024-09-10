package handler

import (
	"Personal_Budget_Tracker/models"
	"Personal_Budget_Tracker/utils"
	"fmt"
	"time"
)

var keywords = []string{
	"new",
	"show",
	"balance",
}
var budget float64
var duration int
var expenseType []models.Money

func CreateBudget() models.BudgetExpense {
	var value string
	fmt.Println("---------------------------------------------------------")
	fmt.Println("******************* CREATING TRACKER ********************")
	fmt.Println("---------------------------------------------------------")

	fmt.Println("Budget Price?")
	fmt.Printf(">> ")
	fmt.Scanln(&budget)

	fmt.Println("Duration? (days)")
	fmt.Printf(">> ")
	fmt.Scanln(&duration)

	fmt.Println("Expense Type?")
	for {
		fmt.Printf(">> ")
		fmt.Scanln(&value)
		if value == "end" {
			break
		}
		expenseType = append(expenseType, models.Money{Title: value, Price: 0.00})
	}
	return models.BudgetExpense{Budget: budget, Expense: expenseType, Duration: utils.DaysToHour(duration), CreatedAt: time.Now()}
}

func DisplayBudget(budget models.BudgetExpense) {
	var total float64
	fmt.Println("************************ DISPLAY ************************")
	fmt.Println("---------------------------------------------------------")

	for _, money := range budget.Expense {
		fmt.Println(money.Title+": ", money.Price)
		total += money.Price
	}
	fmt.Println("Budget: #", budget.Budget, "Total Expenses: #", total, "=> Remaining: #", budget.Budget-total)
	fmt.Println("Start Date:", budget.CreatedAt.Format(time.DateTime), "<==> End date:", budget.CreatedAt.Add(budget.Duration).Format(time.DateTime))
	fmt.Println("*********************************************************")

}

func Init() {
	fmt.Println("*********************************************************")
	fmt.Println(">>>>>>>>>>>>>>> PERSONAL BUDGET TRACKER <<<<<<<<<<<<<<<<<")
	fmt.Println("---------------------------------------------------------")
}
