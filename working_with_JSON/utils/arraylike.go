package utils

import (
	"encoding/json"
	"log"
	"os"
	"working_with_JSON/models"
)

func Reader(student []models.Student) []models.Student {
	//	check if the file is present
	file, err := os.ReadFile("storage/student.json")
	if err != nil {
		log.Fatal(err)
	}
	//file.re
	err = json.Unmarshal(file, &student)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(student)
	return student
}

func Saver(students []models.Student) error {
	jsonType, err := json.Marshal(students)
	if err != nil {
		return err
	}

	// creates if not exists
	file, err := os.Create("storage/student.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonType)
	if err != nil {
		return err
	}
	return nil
}

func StudentAdder(student models.Student) error {
	var tempStudent []models.Student
	students := Reader(tempStudent)
	students = append(students, student)
	err := Saver(students)
	if err != nil {
		return err
	}
	return nil
}
