package main

import "fmt"

func main() {
	//if <boolean expression> { code } else { code }
	x := 27

	if x%2 != 0 {
		fmt.Println(x, "tek sayıdır")
	} else {
		fmt.Println(x, "çift sayıdır")
	}

	myFunc()
}

func myFunc() {
	x := -5
	if x > 0 {
		fmt.Println(x, "negatif sayıdır.")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	} else {
		fmt.Println(x, "tek sayıdır.")
	}
}
