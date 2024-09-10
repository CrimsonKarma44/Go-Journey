package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solution(number int) string {
	numerals := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	stringConv := strings.Split(strconv.Itoa(number), "")
	// separates the number into their separate power level
	func() {
		for i := 0; i < len(stringConv); i++ {
			stringConv[i] = strconv.Itoa(func() int {
				val, _ := strconv.Atoi(stringConv[i])
				return val
			}() * func() int {
				val := 1
				for v := 0; v < len(stringConv)-i-1; v++ {
					val *= 10
				}
				return val
			}())
		}
	}()

	//sorting the keys of the numerals
	var list []int
	func() {
		for i, _ := range numerals {
			list = append(list, i)
		}
		sort.Ints(list)
	}()

	for _, v := range stringConv {
		fmt.Println(v)
		for _, s := range list {
			if val, _ := strconv.Atoi(v); val == 0 {
				fmt.Println(0)
				break
			}
			if val, _ := strconv.Atoi(v); val <= s {
				fmt.Println(s)
				break
			}
		}
	}
	return numerals[number]
}
func main() {
	Solution(440)
}
