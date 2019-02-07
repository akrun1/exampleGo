package main

import (
	"fmt"
)

func main() {

	v1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "arun",
		friends: map[string]int{"james": 32,
			"Missy": 24,
		},
		favDrinks: []string{
			"mango",
			"orange",
		},
	}

	fmt.Println(v1)

	for k, v := range v1.friends {
		fmt.Println(k, v)

	}

	for i, v := range v1.favDrinks {
		fmt.Println(i, v)
	}

}

