package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var aul Customer
	aul.Name = "Aul"
	aul.Address = "Indonesia"
	aul.Age = 22

	fmt.Println(aul)
	fmt.Println(aul.Name)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     25,
	}
	fmt.Println(joko)

	// budi := Customer{"Budi", "Jakarta", 31}
	// fmt.Println(budi)
	joko.sayHi("Eko")
}
