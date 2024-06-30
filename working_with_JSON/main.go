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
	//var date time.Time
	//date = time.Date(2001, time.January, 23, 00, 00, 00, 0, time.Local)
	//student := models.Student{Name: "Princewill", Age: 23, School: "FUTMINNA", DateOfBirth: date}

	//utils.Single(student)

	reader := bufio.NewReader(os.Stdin)
	var person models.Student
	var students []models.Student
	var num int
	var condition string
	for {
		fmt.Print("How many students:")
		fmt.Scan(&num)
		if num == 1 {

			fmt.Println("Name:")
			name, _, _ := reader.ReadLine()
			person.Name = string(name)

			fmt.Println("Age:")
			fmt.Scan(&person.Age)

			fmt.Println("School:")
			school, _, _ := reader.ReadLine()
			person.School = string(school)

			person.DateOfBirth = time.Now()

			err := utils.Single(person)

			if err != nil {
				fmt.Println("Student could not be added!...", err)
			} else {
				fmt.Println("Student Added....")
			}
		} else if num > 1 {

			for i := 0; i < num; i++ {

				fmt.Println("Name:")
				name, _, _ := reader.ReadLine()
				person.Name = string(name)

				fmt.Println("Age:")
				fmt.Scan(&person.Age)

				fmt.Println("School:")
				school, _, _ := reader.ReadLine()
				person.School = string(school)

				// for some reason the dateofbirth refused to save using fmt.Scan
				person.DateOfBirth = time.Now()

				fmt.Println("............")
				fmt.Println("....Next....")
				fmt.Println("............")
				students = append(students, person)
			}
			err := utils.Double(students)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(students)
			}
		}

		fmt.Println("Do you wanna continue? (y/n)")
		fmt.Scan(&condition)
		if condition == "y" {
		} else {
			break
		}
	}
}
