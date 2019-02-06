package main

import (
	"fmt"
)

// We create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS
var x int

type person struct {
	first string
	last  string
}

// not good to use alias without methods
type foo int

var y foo

func main() {
	fmt.Println(x)
}

