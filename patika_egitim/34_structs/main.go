package main

import "fmt"

func main() {
	var employee struct { //structure
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee)
	fmt.Printf("%#v\n", employee)

	employee.name = "Dursun"
	employee.age = 30
	employee.isMarried = true
	fmt.Printf("%#v\n", employee)
	maa()
}

func maa() {
	type employee struct { // underlying type
		name      string
		age       int
		isMarried bool
		kids      []string
	}

	var e1 employee
	e1.name = "Deneme"
	e1.age = 23
	e1.isMarried = true
	fmt.Printf("%#v\n", e1)

	var e2 employee
	e2.name = "Fatma"
	e2.age = 30
	e2.isMarried = false
	fmt.Printf("%#v\n", e2)

	e3 := employee{
		name: "Sabri", age: 5, isMarried: false,
	}
	fmt.Printf("%#v\n", e3)

	e4 := employee{
		name: "Ahmet", age: 25, isMarried: false, kids: []string{"Demet", "Duru"},
	}
	fmt.Printf("%#v\n", e4)
	fmt.Println(e4)
	fmt.Println(e4.kids)
	fmt.Println(e4.kids[0])
}
