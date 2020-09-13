package main

import "fmt"

type geometricFigure interface {
	getArea() float64
}

type triangle struct {
	a float64
	b float64
}
type square struct {
	a float64
}

func main() {
	t := triangle{
		a: 6,
		b: 9,
	}
	printArea(t)
	// fmt.Println(t.getArea())
	s := square{
		a: 10,
	}
	printArea(s)
	// fmt.Println(s.getArea())
}

func printArea(gf geometricFigure) {
	fmt.Println(gf.getArea())
}

func (t triangle) getArea() float64 {
	return (t.a * t.b) / 2
}

func (s square) getArea() float64 {
	return s.a * s.a
}
