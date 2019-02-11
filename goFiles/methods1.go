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

func (s secretAgent) speak(number string) {
	fmt.Println("I am", s.first, s.last, "and id is", number)

}
func main() {
	sa1 := secretAgent{
		person: person{

			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	mp1 := secretAgent{
		person: person{

			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}

	fmt.Println(sa1.person)
	sa1.speak("007")
	mp1.speak("005")

}
