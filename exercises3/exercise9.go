package main

import (
	"fmt"
)

func main() {
	favSport := "Cricket"
	switch favSport {

	case "tennis":
		fmt.Println("I play tennis")
	case "Cricket":
		fmt.Println("England plays cricket")
	default:
		fmt.Println("if nothing prints, this will")
	}

}

