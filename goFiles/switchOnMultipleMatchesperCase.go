package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n{
	case "Moneypenny", "Bond", "Do not" :
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")

	default:
		fmt.Println("this is default")

	}

}

