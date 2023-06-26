package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

// is a relation -- klasik OOP
// has a relation --

func main() {
	e1 := employee{
		name:      "Umut",
		age:       40,
		isMarried: true,
	}
	fmt.Println(e1)

	// e2 := e1
	// fmt.Println(e2) //struct değer paylaşır, referans değil. hafızada farklı yer tutuyorlar
	// e2.name = "Derin"
	// fmt.Println(e2)
	// fmt.Println(e1)

	m1 := manager{
		employee: employee{
			name:      "Ayşe",
			age:       50,
			isMarried: true,
		},
		hasDegree: true,
	}
	fmt.Println(m1)

	//Anonim Struct

	theBoss := struct {
		name  string
		money bool
	}{name: "Boss", money: true}
	fmt.Println(theBoss)

}
