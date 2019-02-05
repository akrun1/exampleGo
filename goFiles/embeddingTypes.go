package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   24,
		},
		first: "something coll",
		ltk:   true,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}
	// inner type is promoted to the outer type
	// if there are collision of inner type with outer type, use the dot notation
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2)
	fmt.Println(p2.age)
}

