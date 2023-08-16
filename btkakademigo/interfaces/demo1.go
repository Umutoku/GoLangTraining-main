package interfaces

type shape interface {
	getArea() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (c circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func school(s shape) float64 {
	return s.getArea()
}

func Demo1() {
	r := rectangle{width: 10, height: 20}
	school(r)
}
