package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Modified() []byte {
	//var student []models.Student
	var content []byte
	f, err := os.Open("storage/student.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	content, err = ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(content)
	//json.Unmarshal(content, &student)
	//fmt.Println(student)
	return content
}
