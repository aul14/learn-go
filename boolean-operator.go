package main

import "fmt"

func main() {
	ujian := 80
	absensi := 75

	// lulusUjian := ujian >= 80
	// lulusAbsensin := absensi >= 80
	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensin)

	// lulus := lulusAbsensin && lulusUjian
	// fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)
}
