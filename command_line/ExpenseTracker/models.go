package ExpenseTracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ExpenseTracker struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListExpense []ExpenseTracker

// Add creates a new todo item and appends it to the list
func (l *ListExpense) Add(description string, amount int) {
	ls := *l
	newId := len(ls) + 1
	newExpense := ExpenseTracker{Id: newId, Description: description, Amount: amount, CreatedAt: time.Now()}
	*l = append(*l, newExpense)
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *ListExpense) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *ListExpense) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// Summation method sums all the expenses,
// and returns the value
func (l *ListExpense) Summation() int {
	amount := 0
	for _, value := range *l {
		amount += value.Amount
	}
	fmt.Println("# Adding...")
	return amount
}

// SumS method sums all the expenses for a specific month,
// and returns the value
func (l *ListExpense) SumS(month int) int {
	amount := 0
	for _, value := range *l {
		if value.CreatedAt.Month() == time.Month(month) {
			amount += value.Amount
		}
	}
	fmt.Fprintf(os.Stdout, "# Adding for %q...\n", time.Month(month))
	return amount
}

// Delete method deletes an Expense item from the list
func (l *ListExpense) Delete(i int) error {
	var ls []ExpenseTracker

	// Adjusting index for 0 based index
	condition := false
	for _, val := range *l {
		if val.Id != i {
			ls = append(ls, val)
		}
		condition = true
	}
	if condition {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = ls
	return nil
}
