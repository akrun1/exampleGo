package main

import (
	"fmt"
)

func main() {
	x := "James Bond"
	if x == "Moneypenny" {
		fmt.Println("Hello, this is ", x)
	} else if x == "James Bond" {
		fmt.Println("Hello, this is ", x)
	} else {
		fmt.Println("Not any of these")
	}
}

