package rangee

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir", "Adana", "Bursa"}

	for i, sehir := range sehirler {
		fmt.Println(i, sehir)
	}
}
