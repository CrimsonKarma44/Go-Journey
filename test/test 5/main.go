package main

import "fmt"

func main() {
	bankInstance := bank{
		name:  "Unity",
		money: money{20000, "Naira"},
	}

	moneyStorage := storage{name: "bank", money: bankInstance.money}
	fmt.Println(bankInstance)
	fmt.Println(moneyStorage)
}

type money struct {
	value    int
	currency string
}

type bank struct {
	name  string
	money money
}

type storage struct {
	money
	name string
}
