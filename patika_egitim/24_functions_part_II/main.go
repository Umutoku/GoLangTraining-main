package main

import (
	"fmt"
	"time"
)

func main() {
	//merhaba2("Umut", 2) // Argument Function Run
	fmt.Println(result(55))

	merhaba3("Berke", 28)

	name := "Er"
	age := 4
	fmt.Println("")
	merhaba3(name, age)

	//farklı pakete ait fonksiyonlara metot denir
	fmt.Println("")
	var x int = 10
	fmt.Println(x)

	var today time.Time = time.Now() // Now ---> method

	fmt.Println(today)
	fmt.Println("")

	// fmt.Print("Lütfen sınav sonucu giriniz: ")
	// reader := bufio.NewReader(os.Stdin)
	// value, _ := reader.ReadString('\n') // _ blank identifier
	// fmt.Println(value)

	fmt.Println("")
	bölüm, kalan := bölme(104, 5)
	fmt.Println(bölüm, kalan)

}

func merhaba2(name string, age int) { // Parametre Function Write
	fmt.Printf("Adım %s, yaşım %d", name, age)
}

func result(grade int) string {

	if grade >= 50 {
		return "geçtiniz"
		// return varsa çıkacak, if koşulu için
	}
	return "kaldınız"
}

func merhaba3(name string, age int) { // Parametre Function Write
	fmt.Printf("Adım %s, yaşım %d", name, age)
}

// 104 / 5 ====> 20 / kalan 4
func bölme(bölünen int, bölen int) (bölüm int, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen
	return bölüm, kalan
}
