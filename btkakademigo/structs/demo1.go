package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Laptop", 1000, 0.1, "Dell"})
}

type product struct {
	name     string
	price    float64
	discount float64
	brand    string
}
