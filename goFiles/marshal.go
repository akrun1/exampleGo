package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{

		First: "James",
		Last:  "Bond",
		Age:   24,
	}

	p2 := person{

		First: "Miss Money",
		Last:  "penny",
		Age:   22,
	}

	people := []person{

		p1,
		p2,
	}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	fmt.Println(string(bs))
	if err != nil {
		fmt.Println(err)
	}
}

