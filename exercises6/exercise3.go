package main

import (
	"fmt"
)

func main() {
	defer f1()
	fmt.Println("Hello playground")

}

func f1() {
	defer func() {
		fmt.Println("arun here")
	}()
	fmt.Println("hello")

}

