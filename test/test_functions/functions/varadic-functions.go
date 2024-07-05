package functions

import "fmt"

func PrincInt(arg ...int) {
	for _, v := range arg {
		fmt.Print(v)
	}
	//fmt.Println(arg)
	return
}
