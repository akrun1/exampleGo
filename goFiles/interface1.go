package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")

}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{

			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	mp1 := secretAgent{
		person: person{

			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}

	fmt.Println(sa1.person)
	sa1.speak()
	mp1.speak()

	fmt.Println(p1.first)
	p1.speak()
	bar(sa1)
	bar(mp1)
	bar(p1)

	//conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

