package main

import (
	"fmt"
)

func main() {
	switch {

	case (1 == 2):
		fmt.Println("should not print")
	case (2 == 2):
		fmt.Println("Yes, this is correct")
	default:
		fmt.Println("if nothing prints, this will")
	}

}

