package interfaces

import "fmt"

func Dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi interface{} = 5
	Dogrula(sayi)
}
