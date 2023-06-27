package main

import "fmt"

func main() {
	x := 50
	x = x - 10 // assignment statement

	x -= 10 // assignment operation
	fmt.Printf("%v, %T", x, x)

	F := -40
	K := float64(F-32)/1.8 + 273
	fmt.Printf("%v, %T", K, K)

	myFunc()
	myFuncOne()
	myFuncTwo()
	myFuncThree()
}

func myFunc() {

	//age :=40
	const myAge = 40
	fmt.Printf("%v, %T", myAge, myAge)

	yaş := 2
	fmt.Println(yaş)

}

const x = 14 // shadowing constlarda da aynı

func myFuncOne() {
	var x = 24
	fmt.Printf("%v, %T", x, x)

}

func myFuncTwo() {
	const D = 4                            //typeless
	const K = 5.4                          //typeless
	fmt.Printf("%v, %T", (D + K), (D + K)) //type conversion gibi K nın veri tipi float64
	fmt.Printf("%v, %T", K, K)             // K veri tipi float64 D veri tipi int

}

func myFuncThree() {
	const x float64 = 6.4
	const y = 4 + x //otomatik olarak float64 oldu
	fmt.Printf("%v, %T", y, y)

}
