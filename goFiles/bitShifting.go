package main

import (
	"fmt"
)



func main() {

	x := 10
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)

	z := x >> 1
	fmt.Println("")
	fmt.Printf("%d\t\t%b\n", z, z)
}

