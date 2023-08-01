package main

import (
	"fmt"
	"learn-go/helper"
)

func main() {
	helper.SayHello("Aul")
	// helper.sayGoodBye("Aul") // error tidak bisa di akses
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error tidak bisa di akses
}
