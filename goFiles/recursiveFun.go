package main

import (
	"fmt"
)

func main() {
	x := factorial(4)
	fmt.Println(x)
	x2 := factorial2(4)
	fmt.Println(x2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func factorial2(n int) int {
	m := 1
	for ; n > 0; n-- {
		m *= n
	}
	return m

}

