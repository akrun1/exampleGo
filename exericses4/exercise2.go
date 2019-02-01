package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 45, 45, 46, 36, 7, 9, 24, 11}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}

