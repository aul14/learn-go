package main

import "fmt"

func main() {
	name := "Rahman"

	if name == "Aul" {
		fmt.Println("Hello Aul")
	} else if name == "Rahman" {
		fmt.Println("Hello Rahman")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hello gais")
	}

	// sebelum penggunaan if short statement
	length := len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
	// sesudah penggunaan if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
