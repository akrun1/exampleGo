package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type foo int

var y foo

const bar = 42

func main() {
	p1 := person{
		first: "Arun",
		last:  "Sas",
	}
	y = bar
	fmt.Println(p1)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)

}

