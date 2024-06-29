package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	C1 := circle{2}
	R1 := rectangle{5, 2}
	//shapes := []shape{C1, R1}
	var justCircle = shape(C1)
	var justRect = shape(R1)
	//for _, shape := range shapes {
	//	fmt.Println(shape.area())
	//}
	fmt.Println(circle{2}.area())
	fmt.Println(justCircle.area())
	fmt.Println()
	fmt.Println(rectangle{5, 2}.area())
	fmt.Println(justRect.area())

	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("this is")
			continue
		}
		fmt.Println(i)

	}
}
