package main

import (
	"fmt"
	"learn-go/database"
	// _ "learn-go/database" // CONTOH BLANK IDENTIFIER yaitu package yg sudah dipanggil tapi tidak digunakan
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
