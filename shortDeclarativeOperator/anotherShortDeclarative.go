package main

import "fmt"

// Declare the VARIABLE "y1"
// assign the value 43
var y1 = 43

// Declare a variable with an identifier "z"
// and the variable with identifier "z" is of type int
// assigns the zero value of type int to "z"
var z int

func main() {

	x := 52
	n, err := fmt.Println("Hello", x, true)
	fmt.Println(n)
	fmt.Println(err)
	y := x + 100
	fmt.Println(y)

	fmt.Println(z)

	fmt.Println(y1)
	foo()
	type T struct { i int; f float64; next *T }
	t := new(T)
	fmt.Println("t.i : ", t.i, "t.f : ", t.f, "t.next : ", t.next)

}

func foo() {
	fmt.Println("Hello", y1)
}
