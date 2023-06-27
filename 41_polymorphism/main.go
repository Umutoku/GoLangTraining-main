package main

import "fmt"

type shape interface {
	area() float64
}

func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println("Alan: ", shape.area())
	}
} //radyatik parametre --> ne kadar shape geleceği belli değil

type triangle struct {
	h float64
	a float64
}

func (t triangle) area() float64 {
	return (t.h * t.a) / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

type rectangle struct {
	a float64
	b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func main() {
	t := triangle{3, 4}
	s := square{4}
	r := rectangle{4, 5}
	printArea(t, s, r)

}
