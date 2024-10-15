package main

import (
	"encoding/json"
	"log"
	"os"
)

func ret() []Data {
	var data []Data
	file, err := os.ReadFile(storagePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
func retStatus(data []Data, status string) []Data {
	return []Data{}
}
