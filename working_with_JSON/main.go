package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"working_with_JSON/models"
	"working_with_JSON/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var person models.Student
	var condition string
	var students []models.Student

	fmt.Println("register student? (y/n)")
	fmt.Scanln(&condition)
	if condition == "y" {

		fmt.Println("Name:")
		name, _, _ := reader.ReadLine()
		person.Name = string(name)

		fmt.Println("Age:")
		fmt.Scan(&person.Age)

		fmt.Println("School:")
		school, _, _ := reader.ReadLine()
		person.School = string(school)

		person.DateOfBirth = time.Now()

		err := utils.StudentAdder(person)

		if err != nil {
			fmt.Println("Student could not be added!...", err)
		} else {
			fmt.Println("Student Added....")
		}
	} else if condition == "n" {
	}

	fmt.Println("Do want to view Students? (y/n)")
	fmt.Scanln(&condition)
	if condition == "y" {
		fmt.Println("The recorded students are:")
		for count, value := range utils.Reader(students) {
			fmt.Println(count+1, value.Name)
		}
	}
}
