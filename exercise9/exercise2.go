package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()

}
func (p *person) speak() {

	fmt.Println("I am", p.first, p.last, ", the person speak")

}

func main() {

	p1 := person{
		first: "Arun",
		last:  "Kirshna",
	}
	p1.speak()
	saySomething(&p1)

}

