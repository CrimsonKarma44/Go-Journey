package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(IsValidIP("123.045.067.089"))
}

func IsValidIP(ip string) bool {
	array := strings.Split(ip, ".")
	if len(array) == 4 {
		if array[len(array)-1] != "" {
			for _, item := range array {
				if string(item[0]) == "0" && len(string(item)) > 1 {
					return false
				}
				v, err := strconv.Atoi(item)
				if err != nil {
					return false
				}
				if v >= 0 && v <= 255 {
					continue
				} else {
					return false
				}
			}
			return true
		}
	}
	return false
}
