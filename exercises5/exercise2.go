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

	x := map[string]person{

		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range x {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, fl := range v.favflavor {
			fmt.Printf("\t%v:\t%v\n", i, fl)
		}
		fmt.Println("--------")

	}
}

