package main

import (
	"fmt"
	"working_with_JSON/utils"
)

func main() {
	//var count int
	var condition int
	//
	//for {
	//	fmt.Print("No: ")
	//	fmt.Scanln(&count)
	//	for i := 0; i < count; i++ {
	//
	//		fmt.Print("Name: ")
	//	}
	//	fmt.Println("Do you want to continue? (y/n)")
	//	fmt.Scanln(&condition)
	//
	//}

	fmt.Println("View? (y/n)")
	fmt.Scanln(&condition)
	if condition == 'y' {
		utils.Modified()
	}
}
