package main

import (
	"fmt"
)

func main() {
	//xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum("james", 2, 4)
	fmt.Println("The total is", x)

}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {

		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total to give the total", sum)
	}

	return sum

}
