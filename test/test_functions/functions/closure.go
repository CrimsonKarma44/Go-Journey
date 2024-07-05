package functions

import "fmt"

// GetModulus Function Closure
func GetModulus() func(int) int {
	count := 0
	return func(x int) int {
		count = count + 1
		fmt.Printf("modulus function called %d times\n", count)
		if x < 0 {
			x = x * -1
		}
		return x
	}
}

type Circle struct {
	Radius float64
}

func State(radius float64) func(increment float64) float64 {
	newRadius := radius
	return func(increment float64) float64 {
		newRadius = increment + newRadius
		return newRadius
	}
}
func (v *Circle) Area(newRadius float64) float64 {
	return newRadius * newRadius
}
