package main

import "fmt"

func main() {
	// cara 1
	person := map[string]string{
		"name":    "Aul",
		"address": "Bekasi",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// cara 2
	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Aul"
	book["ups"] = "salah"
	fmt.Println(len(book))
	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
