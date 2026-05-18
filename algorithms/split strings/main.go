package main

import (
	"fmt"
	"strings"
)

func Solution(str string) []string {
	var vault []string
	for i := 0; i < len(str); i += 2 {
		group := ""
		for v := i; v < i+2; v++ {
			if v < len(str) {
				group += string(str[v])
			}
		}
		if len(group) != 2 {
			group += "_"
		}
		vault = append(vault, group)
	}
	return vault
}

func main() {
	fmt.Println(strings.Split("1.2.3.", "."))
	fmt.Println(Solution("abcd"))
}
