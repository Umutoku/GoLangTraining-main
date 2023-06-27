package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	x := 4
	y := 10
	sum, dif, prod := calculation(x, y)
	println("Toplam: ", sum)
	println("Çıkarma: ", dif)
	println("Çarpma: ", prod)

	println("")
	println("")

	fmt.Print("Lütfen aldığınız notu giriniz: ")
	grade, _ := getGrade()
	var result string
	if grade >= 50 {
		result = "GEçtin."
	} else {
		result = "Kaldın"
	}
	println(result)

	println("")
	println("")

	target := numRand(1, 100)
	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalış")
	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkınız kaldı")
		fmt.Print("Lütfen tahmininizi yapınız: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if num > target {
			fmt.Println("Tahmininiz daha büyük, daha küçük bir sayı giriniz")
		} else if num < target {
			fmt.Println("Tahmininiz daha küçül, daha büyük bir sayı giriniz")
		} else {
			fmt.Println("Tahmininiz doğru", attempts, " seferde buldunuz")
			break
		}

		if attempts == 9 {
			println("Bilemediniz, doğru sayı:", target)
		}
	}

}

func calculation(num1, num2 int) (sum int, dif int, prod int) {
	sum = num1 + num2
	dif = num1 - num2
	prod = num1 * num2

	return sum, dif, prod
}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil
}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
