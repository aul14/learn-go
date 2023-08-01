package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address2 := &address1

	address3 := &address1

	address2.City = "Bandung"

	// address2 = &Address{
	// 	City:     "Malang",
	// 	Province: "Jawa Timur",
	// 	Country:  "Indonesia",
	// }

	*address2 = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)

	address4.City = "Bekasi"
	fmt.Println(address4)
}
