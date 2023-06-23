package main

import (
	"fmt"
	"math"
)

func main() {

	/* 	5
	   	3.14
	   	"passed"
	   	true */

	r := 3.0

	const pi float64 = 3.14 //sabit

	areaOfCircle := pi * (math.Pow(r, 2))

	fmt.Println(areaOfCircle)

	const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	//x = 5, x++, x=x+1 // const değişmez

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

	//const --> compile time
	//var --> run time

	var v = math.Pow(3, 4)
	//const l = math.Pow(3, 2) --> run timeda çalıştığı için kabul edilmez
	fmt.Printf("%T, %v\n", v, v)

	o := 3
	const p = 3
	fmt.Printf("%T, %v\n", o, o)
	fmt.Printf("%T, %v\n", p, p)

	const xx = 1
	var yy = 3
	fmt.Printf("%T, %v\n", (xx + yy), (xx + yy))

	const xxx, yyy = 3, 5
	fmt.Printf("%T, %v\n", (xxx + yyy), (xxx + yyy))

}
