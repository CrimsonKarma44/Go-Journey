package main

import "fmt"

func main() {
	func(l interface{ give() }) {
		l.give()
	}(&fox{})
}

// type animal interface[]
type fox struct{}

func (f *fox) give() {
	fmt.Println("Hello world!")
}
func (f *fox) take() {}
