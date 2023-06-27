package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Practice 1
	x := 75
	var y float64

	y = float64(x) //tpye(value)
	fmt.Println(y)

	myPractice2()
	myPractice3()
	myPractice4()
	myPractice5()
}

// Practice 2
func myPractice2() {
	x := 5
	y := 10
	fmt.Println("X:", x, "Y:", y)
	//Atama işlemi ile x ve y değerleri yer değiştirdi
	x, y = y, x
	fmt.Println("X:", x, "Y:", y)
}

// pratice3 non English variable names -- Kullanımı önerilmez
func myPractice3() {
	yaş := 30
	fmt.Println(yaş)
}

// pratice4 shadowing kavramı, gölgeleme
func myPractice4() {
	x := 5

	if true {
		x := 10 // bu şekilde yeni bir değişken oluşturuyoruz
		fmt.Println(x + 1)
	}

	fmt.Println(x)
}

// practice5 : 40 as a string
func myPractice5() {
	x := 65

	s := string(rune(x))

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", s, s) //"65"

	y := strconv.Itoa(x) // stringe çevirip o şekilde kullanmak için
	fmt.Printf("%v, %T\n", y, y)
}
