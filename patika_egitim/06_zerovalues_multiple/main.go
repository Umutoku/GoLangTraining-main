package main

import "fmt"

func main() {
	var name string // static tpye
	name = "Umut"
	var age int = 21
	torf := true

	var (
		name2 string = "Umut" // static tpye
		age2  int    = 21
		torf2        = true
	)

	// var name3, age3, isMarried3, weight3, height3 = "Umur", 29, false, 187, 93.5
	name3, age3, isMarried3, weight3, height3 := "Umur", 29, false, 187, 93.5
	//zero values --> string = "", int = 0, boolean = false
	fmt.Println(name3, age3, isMarried3, weight3, height3)
	fmt.Println("Hello Malatya", name, age, torf)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", torf)
	fmt.Printf("%T\n", name2)
	fmt.Printf("%T\n", age2)
	fmt.Printf("%T\n", torf2)
}
