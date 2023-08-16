package arrays

import "fmt"

func Demo1() {
	var sayilar [2][3]int

	sayilar[0][0] = 1
	sayilar[0][1] = 2
	sayilar[0][2] = 3
	sayilar[1][0] = 4
	sayilar[1][1] = 5
	sayilar[1][2] = 6

	fmt.Println("")

	for i := 0; i < len(sayilar); i++ {
		for j := 0; j < len(sayilar[1]); j++ {
			fmt.Println(sayilar[i][j])
		}
	}
}
