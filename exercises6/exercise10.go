package main

import (
	"fmt"
)

func main() {
	x1 := foo()
	fmt.Println(x1())
	fmt.Println(x1())
	fmt.Println(x1())
	fmt.Println(x1())
	fmt.Println(x1())
	fmt.Println(x1())

	g1 := foo()
	fmt.Println(g1())
	fmt.Println(g1())
	fmt.Println(g1())
	fmt.Println(g1())
	fmt.Println(g1())
	fmt.Println(g1())

}

func foo() func() int {

	x := 0
	return func() int {
		x++
		return x

	}

}
