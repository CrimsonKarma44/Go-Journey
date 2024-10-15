package ExpenseTracker

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// List method lists all the expenses
func List(arr ListExpense) {
	fmt.Println("# ID   Date       Description  Amount")
	for _, val := range arr {
		fmt.Fprintf(os.Stdout, "# %d", val.Id)
		size := 5 - len(strconv.Itoa(val.Id))
		for i := 0; i < size; i++ {
			fmt.Fprintf(os.Stdout, " ")
		}
		fmt.Fprintf(os.Stdout, "%v-%v-%v  %v", val.CreatedAt.Year(), monthToInt(val.CreatedAt.Month().String()), val.CreatedAt.Day(), val.Description)
		size = 13 - len(val.Description)
		for i := 0; i < size; i++ {
			fmt.Fprintf(os.Stdout, " ")
		}
		fmt.Fprintf(os.Stdout, "$%v\n", val.Amount)
	}
}

func monthToInt(month string) int {
	month = strings.ToLower(month)
	months := map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
	if val, exists := months[month]; exists {
		return val
	}
	return 0
}
