package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments : ", args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	os.Setenv("APP_USERNAME", "root")
	os.Setenv("APP_PASSWORD", "root")

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
