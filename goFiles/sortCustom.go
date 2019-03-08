package main

import (
	"fmt"
	"sort"
)

type Person struct {
	first string
	age   int
}
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

type ByFirst []Person

func (a ByFirst) Len() int {
	return len(a)
}
func (a ByFirst) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByFirst) Less(i, j int) bool {
	return a[i].first < a[j].first
}

func (p Person) String() string {

	return fmt.Sprintf("%s: %d", p.first, p.age)
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 24}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByFirst(people))
	fmt.Println(people)
}
