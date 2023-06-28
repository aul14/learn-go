package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAul NoKTP = "192389928922"
	var marriedStatus Married = true

	fmt.Println(noKtpAul)
	fmt.Println(marriedStatus)
}
