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

	fmt.Println("test", money{20000, "Naira"})
}

type money struct {
	value    int
	currency string
}

func (m money) String() string {
	return fmt.Sprintf("%d in %s", m.value, m.currency)
}

type bank struct {
	name  string
	money money
}

type storage struct {
	money
	name string
}
