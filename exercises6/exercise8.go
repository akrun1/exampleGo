package main

import (
	"fmt"
)

func main() {
	x1 := foo()
	fmt.Println(x1())
}

func foo() func() string {

	return func() string {
		return "hello arun"

	}

}

