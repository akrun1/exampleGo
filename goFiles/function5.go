package main

import (
	"fmt"
)

func main() {
	x := foo(24, 32)
	fmt.Println(x)
}

// func (r receiver) identifier(parameters) (return(s)) {...}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the total which is now", sum)

	}
	fmt.Println("the total is,", sum)
	return sum

}

