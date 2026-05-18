package utils

import (
	"fmt"
	"os"
)

func InitLocalStore(path string) error {
	// Initialize local storage (e.g., create necessary directories)

	pwd, _ := os.Getwd()
	processedPath := pwd + path + "/processed/"
	originalPath := pwd + path + "/original/"

	if os.IsExist(os.MkdirAll(originalPath, os.ModePerm)) {
		return nil
	}
	if os.IsExist(os.MkdirAll(processedPath, os.ModePerm)) {
		return nil
	}

	err := os.MkdirAll(originalPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error creating directory: %s", err)
	}

	err = os.MkdirAll(processedPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error creating directory: %s", err)
	}

	return nil
}
