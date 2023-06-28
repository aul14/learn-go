package main

import "fmt"

func main() {
	// slice adalah reference data array
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei lagi"
	// fmt.Println(months)
	var slice2 = months[10:]
	fmt.Println(slice2)

	//  menggunakan append = membuat slice baru dengan menambah data dari data lama ke posisi terakhir slice,
	//  jika posisi slice penu, maka akan membuat array baru.
	slice3 := append(slice2, "Aul")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// menggunakan make = membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Aul"
	newSlice[1] = "Rahman"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// menggunakan copy = menyalin slice dari souce ke destination
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan penggunaan slice dan array
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
