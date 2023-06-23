//addition + => 15+10 => 15,10 operand + operator

package main

import "fmt"

func main() {
	x, y := 15, 10
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", (x + y), (x + y))
	fmt.Printf("%T, %v\n", (x * y), (x * y))
	fmt.Printf("%T, %v\n", (x / y), (x / y)) //int / int sonuç float değil int çıkacaktır
	fmt.Printf("%T, %v\n", (x % y), (x % y)) //kalan verir

	z := 5 / 2
	fmt.Printf("%T, %v\n", z, z) // sonuç integer çıkar
	k := 5.0 / 2
	fmt.Printf("%T, %v\n", k, k) // sonuç float çıkar

	//Increment ++, Decrement --PostFix
	d := 10

	fmt.Println(d)
	d++
	fmt.Println(d)
	//fmt.Println(d++) //önceden arttırılmalı

	x = x - 1
	fmt.Println(d)

	x--
	fmt.Println(d)

}
