package main

import "fmt"

func main() {
	name := "Kurniawan"

	switch name {
	case "Rahman":
		fmt.Println("Hallo Aul")
		fmt.Println("Salam kenal")
	case "Joko":
		fmt.Println("Hallo Joko")
		fmt.Println("Salam kenal")
	default:
		fmt.Println("Kenalan dong")
		fmt.Println("Boleh kaga?")
	}

	// // sebelum penggunaan switch short statement
	length := len(name)
	// switch length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// // sesudah penggunaan switch short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// switch tanpa kondisi
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
