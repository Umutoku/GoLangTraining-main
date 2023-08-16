package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("saved")
}

func Demo2() {
	c := customer{firstName: "John", lastName: "Smith", age: 25}
	c.save()
}
