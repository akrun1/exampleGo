package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	address string
}

func main() {
	p1 := person{
		first: "arun",
		last:  "kir",
	}
	fmt.Println(p1.first)
	changeMe(&p1)
	fmt.Println(p1.first)

}

func changeMe(p *person) {

	//(*p).first = "keto"
	p.first = "keto"

}

