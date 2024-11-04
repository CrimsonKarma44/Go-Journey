package main

import "net/http"

type test struct {
	name string
}

func (s test) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func main() {
	//http.Handle("/", test)
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
