package main

import (
	"fmt"
)

func main() {
	a1 := foo(5)
	fmt.Println(a1)
	b1 := bar(foo, 10)
	fmt.Println(b1)
}

func foo(xi int) int {

	return xi * xi

}

func bar(f func(xi int) int, ii int) int {

	v1 := f(ii)
	return v1 + 3

}

