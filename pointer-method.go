package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	aul := Man{"Aul"}
	aul.Married()

	fmt.Println(aul.Name)
}
