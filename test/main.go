package main

import (
	"fmt"
	"os"
	"strconv"
)

func viewUsers() []string {
	var profileNames []string

	file, _ := os.ReadDir("../form/profiles")
	for _, f := range file {
		profileNames = append(profileNames, f.Name()[:len(f.Name())-5])
	}
	return profileNames
}

func main() {
	//nameList := viewUsers()
	//fmt.Println(nameList[0])
	//fmt.Println(nameList[1])
	comment := "65"

	value, _ := strconv.Atoi(comment)
	fmt.Printf("%s\n", comment)
	fmt.Printf("%v\n", value)
	//file, _ := os.ReadDir(".")
	//fmt.Println(file)
}
