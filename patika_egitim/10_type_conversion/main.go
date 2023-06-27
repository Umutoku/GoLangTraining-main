package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	//Type Conversion type(value) => int(y)

	fmt.Println(x + int(y))

	fmt.Printf("%v %T\n", y, y) //type conversion yeni bir değer oluşturur. Eski veri tipi değişmez.

	var z int8 = 10
	var k int16 = 10

	fmt.Println(z + int8(k))

	var d int8 = 127
	var e int16

	e = int16(d)

	fmt.Println(e)

	/* 	r := 10
	   	t := "20"

	   	fmt.Printf("%v %t\n", x, x)
	   	fmt.Printf("%v %t\n", y, y)

	   	fmt.Println(x + y) */

	num1 := 106
	str1 := string(rune(num1))

	fmt.Printf("%v %T\n", str1, str1)
}
