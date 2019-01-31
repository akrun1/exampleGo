package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":      32,
		"MoneyPenny": 27,
	}

	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)

	// deleting some key that doesn't exist
	// no errors are thrown
	delete(m, "Ian")
	fmt.Println(m)

	if v, ok := m["MoneyPenny"]; ok {
		fmt.Println(v, " does exist")
		delete(m, "MoneyPenny")

	}
	fmt.Println(m)
}

