package main

import (
	"fmt"
)

func main() {
	x := [5]int{42, 43, 45, 45, 46}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}

