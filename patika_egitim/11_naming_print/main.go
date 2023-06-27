package main

import "fmt"

var k = 10

func main() {
	fmt.Println("merhaba")
	fmt.Print("merhaba")
	fmt.Println("")
	fmt.Printf("merhaba")

	name := "Umut"

	fmt.Print("Benim adım", name)
	fmt.Println("Benim adım", name)
	fmt.Printf("Benim adım %v %T", name, name) //--> Printf kullanılacak ise format eki almalı

	x := 100
	y := 20
	z := 10
	fmt.Printf("%b %d %o", x, y, z) // x ikilik z onluk y sekizlik sistemde gösterecek

	name1, age := "Umur", 29

	fmt.Println("")
	fmt.Println("")
	fmt.Print("Benim adım ", name1, " ve ben ", age, " yaşındayım")
	fmt.Println("")
	fmt.Println("Benim adım", name1, "ve ben", age, "yaşındayım") //println için string boşluk ihtiyacı yoktur
	fmt.Printf("Benim adım %v, ve ben %v yaşındayım", name1, age) //printf bu şekilde formatlayıp sonuç verir

	//VISIBILITY

	//k := 10
	fmt.Println(k)

	myFunc()

	/* 	var coinType string
	   	var custName string
	   	//Go camel case isimlendirme kullanılır
	   	var URL string  // Url değil
	   	var HTTP string // http değil "xyzHTTP" */

	// i, j, k index isimlendirme

}

func myFunc() {
	fmt.Println(k)
}
