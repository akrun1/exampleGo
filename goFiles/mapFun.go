package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Arun"])
	//v, ok := m["Barnabas"]
	//fmt.Println(v, ok)

	// idiomatic way to check if a value is found
	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)

	}
}

