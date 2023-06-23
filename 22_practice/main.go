package main

import "fmt"

func main() {
	myFunc1()
}
func myFunc1() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "çifttir")
		} else {
			fmt.Println(i, "tektir")
		}
	}
}
func myFunc2() { //neredeyse while yapısı
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}
func myFunc3() {
	switch x := 25; {
	case x < 20:
		fmt.Printf("%d küçüktür 20\n", x)
		fallthrough //koşula uysa bile caselere devam edecek
	case x < 100:
		fmt.Printf("%d küçüktür 100\n", x)
		fallthrough
	case x < 200:
		fmt.Printf("%d küçüktür 200\n", x)

	}
}
func myFunc4() { //idiomatic //parçaların daha bağımsız yazılması
	// if i := 20; i%2 == 0 {
	// 	fmt.Println(i, "çifttir")
	// } else {
	// 	fmt.Println(i, "tektir")
	// }
	x := 20
	if x%2 == 0 {
		fmt.Println(x, "çifttir.")
		return
	}
	fmt.Println(x, "tektir")
}
func myFunc5() {
	var x, y int
	for x = 2; x > 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
			if y > (x / y) {
				fmt.Printf("%d bir asal sayıdır\n", x)
			}
		}
	}
}
