package conditionals

import (
	"fmt"
)

func Demo1() {
	//üç adet int değişken tanımlayınız.
	//Ekrana en büyük olan

	a, b, c := 5, 4, 1

	if a > b {
		if a > c {
			fmt.Println("En büyük sayı: ", a)
		} else {
			fmt.Println("En büyük sayı: ", c)
		}
	} else if b > c {
		fmt.Println("En büyük sayı: ", b)
	} else {
		fmt.Println("En büyük sayı: ", c)
	}
}
