package main

import "fmt"

var y string
var z int

func main() {

	// DECLARE a VARIABLE to be of a certain TYPE
	// and then ASSIGN a VALUE of that TYPE to the VARIABLE

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	y = "Arun"
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println("printing the value of z", z, "ending")
	fmt.Printf("%T", z)
}
