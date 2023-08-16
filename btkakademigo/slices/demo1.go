package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)
	fmt.Println(isimler)
	isimler[0] = "Ahmet"
	isimler[1] = "Mehmet"
	isimler[2] = "Ayşe"
	isimler = append(isimler, "Fatma")
	fmt.Println(isimler)
}
