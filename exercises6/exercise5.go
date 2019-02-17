package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {

	return s.length * s.length

}
func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println("circle area is:", s.(circle).area())
	case square:
		fmt.Println("square area is:", s.(square).area())

	}

}

func main() {

	s1 := square{

		length: 52.0,
	}

	c1 := circle{
		radius: 12.0,
	}

	info(s1)
	info(c1)
}

