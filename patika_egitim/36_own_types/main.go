package main

import (
	"fmt"
	"strings"
)

// type employee struct { // struct --> underlying type, employee --> defined type, named type
// 	name      string
// 	age       int
// 	isMarried bool
// }

type mile float64
type kilometer float64
type mystring string

func main() {
	var m1 mile
	m1 = 3.2
	fmt.Printf("%T,%v", m1, m1)

	f1 := float64(4.4)
	fmt.Println(float64(m1) + f1) //m1 mile veri tipi olduğu için float64 yapmadan toplama işlemine almaz
	fmt.Println("")
	fmt.Printf("%T,%v", (m1 + mile(f1)), (m1 + mile(f1)))
	fmt.Println("")
	fmt.Printf("%T,%v", (float64(m1) + f1), (float64(m1) + f1))

	k1 := kilometer(7.8)
	fmt.Println("")
	fmt.Printf("%T,%v", k1, k1)

	m2 := mile(4.6)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(m1 + m2)
	fmt.Println(m1 * m2)
	fmt.Println(m1 + m2 + 2.1)

	fmt.Println("")
	mystring := "Hale"
	fmt.Println(strings.ToUpper(mystring))
	fmt.Println("")

	mEx := mile(10)
	kEx := toKilometer(mEx)
	fmt.Println(kEx)
	fmt.Println("")
	kEx2 := kilometer(10)
	mEx2 := toMile(kEx2)
	fmt.Println(mEx2)
}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {
	return mile(k * 0.62)
}
