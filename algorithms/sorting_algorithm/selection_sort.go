package sorting_algorithm

import "fmt"

// SelectSort is an implementation of the  Selection sort algorithm.
// It is a simple and efficient sorting algorithm that works by repeatedly selecting the smallest
// (or largest) element from the unsorted portion of the list and moving it to the sorted portion of the list.
func SelectSort(arr []int) []int {
	//by default declaration of an int stores 0 (zero)
	var value int
	length := len(arr)

	for i := 0; i < length; i++ {
		minimum := arr[i]
		for j := i + 1; j < length; j++ {
			if minimum > arr[j] {
				minimum = arr[j]
				value = j
			}
		}

		if minimum < arr[i] {
			arr[i], arr[value] = minimum, arr[i]
		}
		fmt.Println(arr)
	}
	return arr
}
