package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"working_with_JSON/models"
)

func Single(student models.Student) error {
	JsonObj, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("storage/student.json", JsonObj, 0777)
	if err != nil {
		//println(err.Error()
		return err
	} else {
		fmt.Println("Student JSON successfully written")
		return nil
	}
}

func Double(student []models.Student) error {
	JsonObj, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("storage/student.json", JsonObj, 0777)
	if err != nil {
		//println(err.Error())
		return err
	} else {
		fmt.Println("Student JSON successfully written")
		return nil
	}
}
