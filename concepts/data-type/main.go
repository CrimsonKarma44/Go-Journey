package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Princewill"
	byteConv := name
	emoji := rune(128577)
	for _, b := range byteConv {
		fmt.Println(string(b), reflect.TypeOf(b))
	}
	fmt.Println(string(emoji), reflect.TypeOf(emoji))

	p := Value{value: "Princewill "}
	p.conv()
	fmt.Println(p.AsciiChr[10])

	v := Value{AsciiChr: []rune{86, 105, 110, 99, 101, 110, 116}}
	v.conv()
	fmt.Println(v.value, p.value)
	test := Value{value: "123456789"}
	test.conv()
	fmt.Println(test.AsciiChr, reflect.TypeOf(func(n int) int {
		return int(test.AsciiChr[n]) - 48
	}(3)))
}
