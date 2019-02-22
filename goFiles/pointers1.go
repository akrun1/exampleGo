package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // gives the address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a // assigning the address to an identifier b pointing (*) to an *int
	fmt.Println(b)

	c := &a
	fmt.Println(c)
	fmt.Println(*c) // * gives the value stored in an address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(*b)

}

