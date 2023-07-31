package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(HasName HasName) {
	fmt.Println("Hello", HasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	eko := Person{
		Name: "Eko",
	}
	SayHello(eko)

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
}
