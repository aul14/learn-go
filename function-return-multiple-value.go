package main

import "fmt"

func getFullName() (string, string, int) {
	return "Aul", "Rahman", 17
}

func main() {
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	// fmt.Println(age)
}
