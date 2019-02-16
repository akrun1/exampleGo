package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak(language string) string {
	x1 := fmt.Sprintf("The person with first name %v and last name %v with age %v speaks %v", p.first, p.last, p.age, language)
	return x1
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   24,
	}
	fmt.Println(p1.first)
	fmt.Println(p1.speak("english"))

}

