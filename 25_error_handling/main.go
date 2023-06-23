package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	result, err := evenNum(11)
	if err != nil { // nil null gibi, eğer ortada hata yoksa dönüş nil olabilir
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğinizi Sayı: ", result)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	result1, err1 := sRoot(16.0)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız", file)
	}

}
func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("Hata, çift sayı giriniz.")
	}
	return num, nil //null değer demek
}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz")
	}
	return math.Sqrt(num), nil
}
