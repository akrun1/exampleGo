package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)
	/*
		x := func() int {
			return 451
		}()
	*/

	x := bar()
	fmt.Printf("%T\n", x)
	i := x()
	fmt.Println(i)
	//fmt.Printf("%T", x)

}

func foo() string {
	return "Hello world"

}

func bar() func() int {
	return func() int {
		return 451
	}

}

