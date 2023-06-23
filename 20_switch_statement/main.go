package main

import "fmt"

func main() {

	grade := 3

	//Switch
	switch grade {
	case 5:
		fmt.Println("Pekiyi")
	case 4: // iki aynı case koşulu olamaz
		fmt.Println("İyi")
		y := 100
		fmt.Println(y) //case içindeki değerler orada kullanılabilir
	case 3:
		fmt.Println("Orta")
	case 2:
		fmt.Println("Kötü")
	case 1:
		fmt.Println("Çok kötü")
	default: //default istediğimiz yere yazabiliriz
		fmt.Println("Geçersiz not")
	}

}
