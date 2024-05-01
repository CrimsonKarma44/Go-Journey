package helper

import (
	"fmt"
	"strings"
)

func GetUserInput(remainingTickets uint) (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Println("What is your last name?")
		fmt.Scan(&lastName)
		if len(lastName) > 2 {
			break
		} else {
			fmt.Println("Too short try again!")
		}
	}

	for {
		fmt.Println("First name?")
		fmt.Scan(&firstName)
		if len(firstName) > 2 {
			break
		} else {
			fmt.Println("Too short try again!")
		}
	}

	for {
		fmt.Println("Email?")
		fmt.Scan(&email)
		if strings.Contains(email, "@") && len(email) > 4 {
			break
		} else {
			fmt.Println("Invalid Email!")
		}
	}

	for {
		fmt.Printf("Ticket: ")
		fmt.Scan(&userTickets)
		if userTickets <= remainingTickets {
			break
		} else {
			fmt.Println("Tickets over limit!!")
			fmt.Printf("Only %v tickets remaining\n", remainingTickets)
		}
	}
	return firstName, lastName, email, userTickets
}
