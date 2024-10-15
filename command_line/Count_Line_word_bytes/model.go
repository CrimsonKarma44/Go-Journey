package main

import (
	"encoding/json"
	"os"
	"time"
)

type Data struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

func (d *Data) save() error {
	data := ret()
	data = append(data, *d)

	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	file, err := os.Create(storagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(res)
	if err != nil {
		return err
	}
	return nil
}
