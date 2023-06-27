package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//for <init statement>; <condition>; <post statement> { ---------- }

	myFunc()
}

func myFunc() {
	// i := 0
	// for ; i <= 10; i++ { //i tanımı burada belirtilmiş gibi
	// 	fmt.Println(i)
	// }

	// for { //Infinite Loop
	// 	fmt.Println("Sonsuz döngü")
	// }

	i := 10

	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
func myFunc2() {
	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			continue //bu koşulu sağlarsan for döngüsü başına gider
		}
		fmt.Println(i)
	}
}

func myFunc3() {
	for i := 0; i <= 10; i++ {
		if i == 3 {
			break //bu koşulu sağlarsan döngüden çıkar
		}
		fmt.Println(i)
	}
}
