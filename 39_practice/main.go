package main

import (
	"fmt"
)

type myType int

type rectangle struct {
	a int
	b int
}

type family struct {
	name      string
	age       int
	isMarried bool
}

func main() {

	// 1. Soru
	x := myType(10)
	fmt.Printf("%T,%v", x, x)
	fmt.Println("")
	fmt.Println("")

	// 2. Soru
	y := 10
	z := &y
	*z = 20
	fmt.Println(y)
	fmt.Println("")
	fmt.Println("")

	// 3. Soru

	r1 := rectangle{3, 8}
	r1.display()
	fmt.Println("Alan: ", r1.area())
	fmt.Println("Çevre: ", r1.circumference())
	fmt.Println("")
	fmt.Println("")

	// 4. Soru
	families := showFamily()
	for i := 0; i < len(families); i++ {
		fmt.Printf("%v,%T\n", families[i], families[i])
	}

}

func showFamily() []family {
	family1 := family{
		name:      "Bulut",
		age:       25,
		isMarried: false,
	}

	family2 := family{
		name: "Derme",
		age:  14,
	}
	family3 := family{
		"Umut",
		30,
		true,
	}
	var family4 family
	family4.name = "Buse"
	family4.age = 36
	family4.isMarried = true

	return []family{family1, family2, family3, family4}
}

func (r rectangle) display() {
	fmt.Println("Kenar uzunlukları: ", r.a, " ve ", r.b, " olan dikdörtgen")
}

func (r rectangle) area() int {
	return r.a * r.b
}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}
