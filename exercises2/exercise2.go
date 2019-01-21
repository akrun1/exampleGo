package main

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (24 <= 25)
	c := (49 >= 5)
	d := (5 != 4)
	e := (16 < 22)
	f := (15 > 4)
	fmt.Println(a, b, c, d, e, f)
}
