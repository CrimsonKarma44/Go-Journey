package main

import (
	"fmt"
	"test_interface/functions"
)

func main() {
	value := functions.GetArea()
	functions.View(2, 3, value)

	// function closure use case
	modulus := functions.GetModulus()
	fmt.Println(modulus(-1))
	modulus(2)
	fmt.Println(modulus(-5))
	fmt.Println(modulus(-5))

	//	test
	innerCircle := functions.Circle{Radius: 32}
	newRadius := functions.State(innerCircle.Radius)

	fmt.Println("new radius: ", newRadius(23))
	fmt.Println("area: ", innerCircle.Area(newRadius(23)))
	fmt.Println("area: ", innerCircle.Area(newRadius(23)))

	// Varadic function
	functions.PrincInt(1, 2, 3, 4, 5)
}
