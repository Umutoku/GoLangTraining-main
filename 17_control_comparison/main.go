package main

import "fmt"

func main() {
	//Comparison Operators
	x, y := "a", "b" // string ifadeler ASCII olarak numerik
	//z, q := 3, 5.0   //farklı veri tipleri karşılaştırılamaz
	fmt.Printf("%T, %v\n", x == y, x == y)
	fmt.Printf("%T, %v\n", x < y, x < y)
	fmt.Printf("%T, %v\n", x != y, x != y)
	fmt.Printf("%T, %v\n", x > y, x > y)

	//fmt.Printf("%T, %v\n", z > q, z > q)

	myMyFunc()
}

func myMyFunc() {
	//Logical Operators
	x, y := 15, 20
	set1 := (x == y) //false
	set2 := (x < y)  //true

	fmt.Println("")
	fmt.Printf("%T, %v\n", set1, set1)
	fmt.Printf("%T, %v\n", set2, set2)

	fmt.Printf("%T, %v\n", (set1 && set2), (set1 && set2)) // false && true --> false
	fmt.Printf("%T, %v\n", (set1 || set2), (set1 || set2)) // false || true --> true
}
