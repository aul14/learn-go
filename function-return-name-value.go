package main

import "fmt"

func getFullName2() (firstName string, lastName string, age int) {
	firstName = "Aul"
	lastName = "Rahman"
	age = 17

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
