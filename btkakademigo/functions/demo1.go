package functions

import (
	"fmt"
)

func Topla(sayi1 int, sayi2 int) int {
	return sayi1 + sayi2
}

func Cikar(sayi1, sayi2 int) int {
	return sayi1 - sayi2
}

func Carp(sayi1, sayi2 int) int {
	return sayi1 * sayi2
}

func Bol(sayi1, sayi2 float64) float64 {
	return sayi1 / sayi2
}

func SelamVer() {
	fmt.Print("Merhaba")
}
