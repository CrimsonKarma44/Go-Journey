package main

import (
	"fmt"
	"hashFunction/hash"
)

func main() {
	value := hash.HashFunc{}.New()
	fmt.Println(value)
	fmt.Println("Hello, World!")
}
