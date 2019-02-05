package main

import (
	"fmt"
)

var x, y int

func main() {
	x = 42
	y = 26
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

}

