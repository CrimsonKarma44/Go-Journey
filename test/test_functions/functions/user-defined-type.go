package functions

import "fmt"

// function as User defined type
type Area func(a, b int) float64

func GetArea() Area {
	return func(a, b int) float64 {
		value := a * b
		return float64(value)
	}
}
func View(x, y int, a Area) {
	fmt.Printf("Area: %f\n", a(x, y))
}
