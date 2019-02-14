package main

import "fmt"

func main() {
	v := []int{1, 3, 5, 7, 9}
	fmt.Println(foo(v...))
	fmt.Println(bar(v))
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total

}

