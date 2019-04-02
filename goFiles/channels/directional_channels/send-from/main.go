package main

import "fmt"

func main() {

	cs := make(chan<- int)
	cr := make(<-chan int)
	c := make(chan int)

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	cr = c
	fmt.Println(cr)
	fmt.Printf("c\t%T\n", (<-chan int)(c))
}
