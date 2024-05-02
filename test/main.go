package main

import (
	"encoding/json"
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

// Define a struct type
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
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

	// Create instances of Person
	person1 := Person{Name: "Alice", Age: 30, Country: "USA"}
	person2 := Person{Name: "Bob", Age: 35, Country: "Canada"}

	// Append the struct instances to an array
	people := []Person{person1, person2}

	// Marshal the array into JSON
	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}
