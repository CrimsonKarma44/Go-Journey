package main

import (
	"algorithms/sorting_algorithm"
	"fmt"
)

func main() {
	// fmt.Println(sorting_algorithm.BubbleSort([]int{10, 4, 5, 6, 7}))
	// fmt.Println(sorting_algorithm.SelectSort([]int{64, 25, 12, 22, 11}))
	fmt.Println(sorting_algorithm.InsertionSort([]int{23, 1, 10, 5, 2}))
}
