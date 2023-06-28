package main

import "fmt"

func main() {
	var name string

	name = "Aulia Rahman"
	fmt.Println(name)

	name = "Aul aja"
	fmt.Println(name)

	var friendName = "Budiono"
	fmt.Println(friendName)

	var age = 17
	fmt.Println(age)

	// deklarasi awal lebih singkat tidak menggunakan var
	country := "Indonesia"
	fmt.Println(country)

	country = "Amerika"
	fmt.Println(country)

	// deklarasi multiple variable
	var (
		firstName = "Aul"
		lastName  = "Rahman Test"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
