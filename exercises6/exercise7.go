package main

import (
	"fmt"
)

var g func() = func() {
	fmt.Println("g from outside")
}

func main() {
	v1 := func(x1 int, x2 int) int {
		return x1 + x2
	}
	g()
	fmt.Println(v1(12, 24))
	fmt.Printf("%T\n", v1)

}

