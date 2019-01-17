package main

import "fmt"

var y = 42

// DECLARE the VARIABLE with IDENTIFIER "z"
// is of TYPE string
// and ASSIGN the VALUE "Shaken not"
// It is a STATIC programming language
// VARIABLE is DECLARED to hold a VALUE of certain TYPE

var z string = "Shaken not"
var a string = `James said, "Hello 

     not"`
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)


}

