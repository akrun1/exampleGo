package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 3, 5, 7}
	s := sum(ii...)
	fmt.Println(s)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {

		total += v
	}
	return total

}

