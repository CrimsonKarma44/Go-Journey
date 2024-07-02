package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"working_with_JSON/models"
)

func jsonAdder(student []models.Student) {
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
	fmt.Println(student)
}
