package main

import "fmt"

func main() {
	var sampleInt any
	//var sampleString string

	sampleInt = 32

	switch sampleInt.(type) {
	case int:
		fmt.Println("just what am looking for Interger")
	case string:
		fmt.Println("yep a string alright")
	default:
		fmt.Println("just what no one is expecting")
	}
}
