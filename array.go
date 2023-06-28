package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Aul"
	names[1] = "Rahman"
	names[2] = "Aja"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		75,
	}
	fmt.Println(values)

	// fungsi len untuk menghitung panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))
}
