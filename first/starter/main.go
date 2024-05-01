package main

import (
	"fmt"
	// "strings"
	// "strconv"
	"first/helper"
	"sync"
	"time"
	// provides basic synchronization functionality
)

// Package level variable
const conferenceTickets uint = 50

var remainingTickets uint = 50
var conferenceName string = "Go Conference"

// var bookings = []string{}
// array
// var bookings = [50]string{}
// slices
// list of maps
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// lets one to print type of any variable
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUsers()

	// can run indefinitely without a condition
	// for i := 0; i < count; i++ {
	// for {
	// a for loop in go is similar ot a while loop and a for loop in C all merged in none

	firstName, lastName, email, userTickets := helper.GetUserInput(remainingTickets)

	bookings = bookTickets(firstName, lastName, email, userTickets)
	wg.Add(1)
	// sets the number of goroutine to wait for (increase the counter by the provided number)
	go sendTickets(userTickets, firstName, lastName, email)

	firstNames := getFirstNames()
	fmt.Printf("The first names of the bookings are: %v\n", firstNames)

	var noTicketsRemaining bool = remainingTickets == 0

	if noTicketsRemaining {
		fmt.Println("Our program is booked out. Come back next year.")
		// break
	}

	wg.Wait()
	// blocks until the WaitGroup counter is 0

	// city := "London"

	// switch city {
	// case "New Yock":
	// 	// some code here
	// case "Singapore":
	// 	// some code here
	// case "London", "Berlin":
	// 	// some code here for london & berlin
	// case "Mexico City":
	// 	// some code here
	// case "Hong Kong":
	// 	// some code here
	// default:
	//  // fmt.Println("No city selected")
	// }

}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avaialable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	// _this is known as a blank identifier used when a variable needs to be ignored
	// similar to "for index, value in enumerate(array):" in python
	// for _, booking := range bookings {
	// 	var names = strings.Fields(booking)
	// 	firstNames = append(firstNames, names[0])
	// }
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstname"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// func bookTickets(firstName string, lastName string, email string, userTickets uint) []map[string]string {
func bookTickets(firstName string, lastName string, email string, userTickets uint) []UserData {
	remainingTickets = remainingTickets - userTickets

	// var userData = make(map[string]string)
	// userData["firstname"] = firstName
	// userData["lastname"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	// bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("List of bookings: %v\n	", bookings)

	fmt.Printf("Thank you %v %v for bookign %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

	return bookings
}
func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	// means 10 sec
	time.Sleep(10 * time.Second)
	// for storing a formated string print in a variable
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket\n %v \nto email address %v\n", tickets, email)
	fmt.Println("###################")
	wg.Done()
	// this decrements the waitgroup by 1 and is called a goroutine to indicate that it's finished
}
