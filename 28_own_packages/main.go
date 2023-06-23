package main

import (
	"fmt"
	"getcelcius"
)

func main() {

	// merhaba.Hola()
	// merhaba.Hello()
	// merhaba.Merhaba()

	// fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	// celcius, err := getcelcius.GetCelcius()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// kelvin := celcius + 273
	// fmt.Println(celcius, "Celcius derecesi ", kelvin, "Kelvin derecesine eşittir.")

	fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	celcius, err := getcelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}
	fahrenheit := (celcius * 9 / 5) + 32
	fmt.Println(celcius, "Celcius derecesi ", fahrenheit, "Fahrenheit derecesine eşittir.")
}
