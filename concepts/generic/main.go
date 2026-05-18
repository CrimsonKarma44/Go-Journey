package main

import "fmt"

func main() {
	f := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	i := []int64{1, 2, 3, 4, 5}

	s1 := genericSum(f)
	s2 := genericSum2(i)

	fmt.Println("Sum for float64 :", s1)
	fmt.Println("Sum for int64 :", s2)

}

func genericSum[N int64 | float64](nums []N) N {
	var sum N

	for _, num := range nums {
		sum += num
	}

	return sum
}

type Number interface {
	int64 | float64
}

func genericSum2[N Number](nums []N) N {
	var sum N

	for _, num := range nums {
		sum += num
	}

	return sum
}
