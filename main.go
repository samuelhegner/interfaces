package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	c := circle{
		radius: 10,
	}

	r := rect{
		width:  10,
		height: 23,
	}

	printArea([]geometry{r, c})
	printPerim(c)
	printPerim(r)
}

func printArea(gs []geometry) {

	for _, g := range gs {
		t := reflect.TypeOf(g)
		fmt.Println(t, "area:", g.area())
	}
}

func printPerim(g geometry) {
	t := reflect.TypeOf(g)
	fmt.Println(t, "perim:", g.perim())
}

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
