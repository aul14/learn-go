package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aul Rahman", "Aul"))
	fmt.Println(strings.Split("Aul Rahman Aja", " "))
	fmt.Println(strings.ToLower("Aul Rahman Aja"))
	fmt.Println(strings.ToUpper("Aul Rahman Aja"))
	fmt.Println(strings.ToTitle("aul rahman aja"))
	fmt.Println(strings.Trim("       aul rahman aja      ", " "))
	fmt.Println(strings.ReplaceAll("aul rahman aja aul", "aul", "budi"))

}
