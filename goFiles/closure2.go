package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
}

func foo() {
	fmt.Println("hello")

}

