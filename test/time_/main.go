package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	arg, err :=
		func() (string, error) {
			if len(os.Args) == 2 {
				return os.Args[1], nil
			}
			return "", fmt.Errorf("argument mismatch")
		}()
	if err != nil {
		panic(err)
	}
	book := Book{}.New(arg)
	fmt.Println(book.IsPresent("vishal"))
}

type Book struct {
	path    string
	content *os.File
}

func (b Book) New(path string) Book {
	suffixPath, _ := os.Getwd()
	fullPath := suffixPath + "/" + path
	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	return Book{path, file}
}

func (b *Book) IsPresent(query string) bool {
	file, err := os.ReadFile(b.path)
	if err != nil {
		panic(err)
	}
	buffer := bufio.NewReader(bytes.NewBuffer(file))
	count := 0
	for {
		content, err := buffer.ReadString('\n')
		if content == "\n" {
			break
		}
		if err == io.EOF {
			break
		}
		count++
		newContent := strings.Split(strings.Split(content, "\n")[0], " ")
		for _, value := range newContent {
			if value == query {
				return true
			}
		}
	}
	return false

}
