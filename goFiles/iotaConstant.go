package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	g = iota

)

const (
	d = iota
	e
	f

)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(g)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)


}
