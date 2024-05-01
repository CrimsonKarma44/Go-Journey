package main

import (
	"fmt"
)

func stray(love *int) {
	*love = 1000
}

type Person struct {
	name               string
	email              string
	age                uint
	educational_status string
}

func main() {
	x := 0xFF
	y := 0x9C
	// another way to declear pointer
	var z *int = &x
	// there is no need to specify the datatype for a pointer

	var chidi Person = Person{"Princewill", "vincentprincewill44@gmail.com", 23, "undergratuate"}

	people := &chidi

	fmt.Println(people.name)
	fmt.Println((*people).email)
	people.name = "Cindy"
	fmt.Println(people.name)
	people.email = "Cindylove@gmail.com"
	fmt.Println((*people).email)

	// Displaying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)
	fmt.Printf("address of x in decimal is %v\n", z)

	*z = 20
	fmt.Println("")
	fmt.Printf("changed Value of x is %v\n", x)

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)

	stray(&x)
	fmt.Printf("New value of x is %v\n", x)
}
