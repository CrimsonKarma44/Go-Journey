package main

import "fmt"

func main() {
	nums := make([]int, 0, 2) // Start with capacity 2
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Printf("After appending %d: len=%d, cap=%d\n", i, len(nums), cap(nums))
	}
}
