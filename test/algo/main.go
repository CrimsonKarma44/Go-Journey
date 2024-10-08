package main

import "fmt"

func main() {
	testValue := "{{}}()<"
	value := validator(testValue)
	fmt.Println(value)
}

func validator(item string) bool {
	//defaultValue := make(map[string]int)
	defaultValue := map[string]int{
		"{": 0,
		"}": 0,
		"<": 0,
		">": 0,
		"(": 0,
		")": 0,
		"[": 0,
		"]": 0,
	}

	for _, i := range item {
		defaultValue[string(i)] += 1
	}

	if defaultValue["{"] == defaultValue["{"] {
		if defaultValue["("] == defaultValue[")"] {
			if defaultValue["<"] == defaultValue[">"] {
				if defaultValue["["] == defaultValue["]"] {
					return true
				}
			}
		}
	}
	return false
}
