package main

import (
	"fmt"
	"strings"
)

func main() {
	var value = "love"
	var singleValue = ".map"
	reader := make(map[string]any)
	for i := range strings.SplitSeq(value, "") {
		reader[i] = ""
	}
	if _, ok := reader["."]; ok {
		fmt.Println("it has a file extention")
	}
	reading := strings.Split(value, ".")
	secondReading := strings.Split(singleValue, ".")

	fmt.Println(len(secondReading))
	fmt.Println(secondReading)
	fmt.Println(len(reading))
	fmt.Println(reading)
}
