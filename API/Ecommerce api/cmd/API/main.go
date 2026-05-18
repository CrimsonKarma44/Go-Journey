package main

import (
	"ecommerce"
	"ecommerce/urls"
)

// main this is the terminal by which the program is accessed
func main() {
	urls.Urls(ecommerce.Init())
}
