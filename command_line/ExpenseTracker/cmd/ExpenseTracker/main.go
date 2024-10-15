package main

import (
	"ExpenseTracker"
	"fmt"
	"os"
	"strconv"
	"time"
)

const filename = ".expense.json"

func main() {
	ls := ExpenseTracker.ListExpense{}
	ls.Get(filename)
	if len(os.Args) > 1 {
		switch {
		case os.Args[1] == "summary":
			if len(os.Args) == 4 && os.Args[2] == "--month" {
				monthArg, _ := strconv.Atoi(os.Args[3])
				monthVal := time.Month(monthArg)
				sum := ls.SumS(monthArg)
				fmt.Fprintf(os.Stdout, "# Total expenses for %v : $%d\n", monthVal, sum)
			} else {
				sum := ls.Summation()
				fmt.Fprintln(os.Stdout, "# Total expenses: $"+strconv.Itoa(sum))
			}
		case os.Args[1] == "add":
			var description string
			var amount int
			var err error
			if len(os.Args) == 6 {
				if os.Args[2] == "--description" {
					description = os.Args[3]
				} else {
					description = os.Args[5]
				}

				if os.Args[2] == "--amount" {
					amount, err = strconv.Atoi(os.Args[3])
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					}
				} else {
					amount, err = strconv.Atoi(os.Args[5])
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					}
				}
				if err = ls.Get(filename); err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				ls.Add(description, amount)

				if err = ls.Save(filename); err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}

				fmt.Fprintf(os.Stdout, "# Expense added successfully (ID: %v)\n", ls[(len(ls)-1)].Id)
			}
		case os.Args[1] == "list":
			ls.Get(filename)
			ExpenseTracker.List(ls)
		case os.Args[1] == "delete":
			ls.Get(filename)
			if len(os.Args) == 4 && os.Args[2] == "--id" {
				val, _ := strconv.Atoi(os.Args[3])
				err := ls.Delete(val)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				ls.Save(filename)
				fmt.Println("# Expense deleted successfully")
			}
		}
	}
}
