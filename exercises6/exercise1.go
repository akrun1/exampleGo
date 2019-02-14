package main

import "fmt"

func main() {
	n := foo()
	x, s := bar()
	fmt.Println(n)
	fmt.Println(x, s)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 10, "arun"
}

