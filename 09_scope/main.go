package main

import "fmt"

var packVar = "Package Scope" // paket değişkenlerde kısa gösterim yapamayız
var funcVar = "Func(Package) Scope"

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	var name = "Mehmet"

	name, surname := "Ahmet", "Demi" //redeclaration ile tekrar declare edilebilir

	fmt.Println(name, surname)

	var funcVar = "Func Scope"

	fmt.Println(funcVar)
	fmt.Println(packVar)

	myFunc()
}

func myFunc() {
	fmt.Println(funcVar)
}
