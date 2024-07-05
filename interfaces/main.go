package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c Circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func main() {
	shapes := []Shape{&Rectangle{10, 10}, &Circle{10}}
	for _, shape := range shapes {
		fmt.Println(shape.area())
		fmt.Println(shape.perimeter())
	}
}
