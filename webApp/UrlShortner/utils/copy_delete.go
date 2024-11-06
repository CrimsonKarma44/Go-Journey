package utils

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func CopyUrl(url string) error {
	// Copy the text to the clipboard
	err := clipboard.WriteAll(url)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
		return err
	}
	return nil
}

//func DeleteUrl(path string) error {}
