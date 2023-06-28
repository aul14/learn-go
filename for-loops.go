package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke - ", counter)
		counter++
	}

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke - ", counter)
	}

	// sebelum menggunakan for range
	slice := []string{"Aul", "Rahman", "Aja", "Joko", "Budi"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// menggunakan for range dengan key index
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// menggunakan for range tidak dengan key index
	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Aul"
	person["title"] = "Programmer"

	for _, value := range person {
		fmt.Println(value)
	}
}
