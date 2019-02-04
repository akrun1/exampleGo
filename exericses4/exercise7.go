package main

import (
	"fmt"
)

func main() {

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Hellooooo, James."}

	fmt.Println(xs1)
	fmt.Println(xs2)

	x := [][]string{xs1, xs2}

	for i, xs := range x {
		fmt.Println("record:", i)
		for j, val := range xs {
			fmt.Printf("\tindexposition: %v \t value: \t %v \n", j, val)
		}
	}

}
