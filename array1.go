package main

import (
	"fmt"
)

func main() {
	var x [5]int
	var y [5]*float64

	fmt.Println(x)
	x[4] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(y)
}

