package main

import "fmt"

func main() {

	x := 10

	if x := -5; x > 0 { //noktalı virgül ile ayırabiliriz ama if dışında çalışmaz
		fmt.Println(x, "negatif sayıdır.")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	} else {
		fmt.Println(x, "tek sayıdır.")
	}

	fmt.Println(x) // if dışındaki if değerini alır
}

func myFunc() {
	if x := 25; x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "çift sayıdır")
		} else {
			fmt.Println(x, "tek sayıdır")
		}
	}
}
