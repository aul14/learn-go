package main

import "fmt"

func main() {
	name := "Aul"
	counter := 0

	increment := func() {
		name := "Rahman"
		fmt.Println("INCREMENT")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
