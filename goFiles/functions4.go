package main

import (
	"fmt"
)

func main() {

	foo()
	bar("arun")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")

}

// when we create a function, it is the parameters
// call the function, it would be arguments

// everything in GO is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)

}

// return string
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)

}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b

}
