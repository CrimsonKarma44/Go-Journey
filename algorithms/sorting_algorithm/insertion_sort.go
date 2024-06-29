package sorting_algorithm

import "fmt"

// InsertionSort is a simple sorting algorithm that works by iteratively inserting each element of an unsorted
// list into its correct position in a sorted portion of the list. It is a stable sorting algorithm, meaning
// that elements with equal values maintain their relative order in the sorted output.
func InsertionSort(data []int) []int {
	//var temp int
	//var conditon bool
	index := len(data) - 1
	for i := 0; i < len(data); i++ {
		count := len(data) - i - 1
		for j := i - 1; j >= 0; j-- {
			fmt.Println(data[i], data[j])
			if data[i] < data[j] {
				index = count
				fmt.Println(index, ",", count)
			} else {
				break
			}
			count--
		}
		//data[j] = data[i]
		//data[i] = temp
		//temp = data[i+1]
		fmt.Println()
	}
	return data
}
