package loops

import "fmt"

func FriendsNumber() {
	sayi1 := 220
	sayi2 := 284
	total1 := 0
	total2 := 0
	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			total1 += i
		}
	}
	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			total2 += i
		}
	}
	fmt.Println(total1, total2)
	if total1 == sayi2 && total2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayılardır", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayılar değildir", sayi1, sayi2)
	}
}
