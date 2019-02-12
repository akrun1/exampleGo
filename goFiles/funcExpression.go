package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	f := func() {
		fmt.Println("my first func expression")
	}

	f1 := func(x int) {
		fmt.Println("the age of inncocence", x)
	}

	f()
	f1(20)
}

