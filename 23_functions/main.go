package main

import "fmt"

func main() {
	var x, y, sum int

	x = 5
	y = 10
	sum = x + y

	fmt.Printf("%d ve %d toplamı %d", x, y, sum)

	x = 10
	y = 20
	sum = x + y //eğer burada belirtmez isek en son sum üzerindeki değerden toplar
	fmt.Printf("%d ve %d toplamı %d", x, y, sum)

	//Moduler Programming
	//Readable code
	//From Complex to Simple
	fmt.Println("")
	fmt.Println("")
	fmt.Println(sumMy(5, 10))
	fmt.Println("")
	merhaba()
	fmt.Println("")

	//return vs print
	z := sumMy(10, 5)
	sum2(5, 15)
	fmt.Println(z)       //
	merhaba2("umut", 29) // Sıralama düzgün şekilde verilmeli

	//Fonksiyon isimlendirme
	//İlk karakter harf olacak --> fonksiyon
	//camel Case -- mySum, myBestFunction
	//paket dışında kullanılacaksa ilk harf büyük olmalı
}

func sumMy(x, y int) int {
	return x + y
}
func merhaba() {
	fmt.Println("Benim adım ne?")
}

func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d", name, age)
}

func sum2(x, y int) { //toplamı yazdırıyor
	fmt.Println(x + y)
}
