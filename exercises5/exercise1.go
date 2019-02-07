package main

import (
	"fmt"
)

type person struct {
	first     string
	last      string
	favflavor []string
}

func main() {
	p1 := person{
		first:     "arun",
		last:      "kir",
		favflavor: []string{"vanilla", "chocolate"},
	}
	p2 := person{
		first:     "james",
		last:      "bond",
		favflavor: []string{"chocolate", "mango"},
	}

	x := append(p1.favflavor, p2.favflavor...)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favflavor {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favflavor {
		fmt.Println(i, v)
	}

	for _, v := range x {
		fmt.Println(v)
	}
}

