package main

import "fmt"

var y = 42
var z = false

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%v\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%o\n", y)
	fmt.Printf("%#o\n", y)
	fmt.Printf("%o\t%b\t%#o\n", y, y, y)
	fmt.Printf("%+q\n", 42)

	// Sprint returns the resulting string
	// and can be assigned to a VARIABLE

	y1 := fmt.Sprintln(y)
	y2 := fmt.Sprintf("%o\t%b\t%#o\n", y, y, y)
	fmt.Printf(y1)
	fmt.Printf(y2)


}

