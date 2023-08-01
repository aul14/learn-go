package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Aul")
	data.PushBack("Rahman")
	data.PushBack("Aja")
	data.PushFront("Budi")

	// MENAMPILKAN DATA DARI DEPAN KE BELAKANG
	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	// MENAMPILKAN DATA DARI BELAKANG KE DEPAN
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
