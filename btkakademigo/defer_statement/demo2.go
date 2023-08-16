package defer_statement

import "fmt"

func Demo2(sayi int) string {
	if sayi%2 == 0 {
		return "cift"
	}
	if sayi%2 == 1 {
		return "tek"
	}
	defer DeferFunc()
	return "hata"
}

func Test() {
	fmt.Println(Demo2(11))
}

func DeferFunc() {
	fmt.Println("Bitti")
}
