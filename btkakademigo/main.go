package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/channels"
	"golesson/conditionals"
	"golesson/defer_statement"
	"golesson/error_handling"
	"golesson/functions"
	"golesson/interfaces"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	"golesson/project"
	"golesson/rangee"
	"golesson/restful"
	"golesson/slices"
	"golesson/string_functions"
	"golesson/structs"
)

func main() {

	conditionals.Demo1()
	loops.Demo1(10)
	loops.FriendsNumber()
	arrays.Demo1()
	slices.Demo1()
	slices.Demo2()
	fmt.Println(functions.Topla(1, 2))
	fmt.Println(functions.DortIslem(6, 3))
	fmt.Println(functions.ToplaVariadic(6, 3, 5, 7, 3, 4))

	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11, 11, 11}
	fmt.Println(functions.ToplaVariadic(sayilar...))

	maps.Demo1()

	rangee.Demo1()
	rangee.Demo2()
	rangee.Demo3()

	sayi := 20
	pointers.Demo1(&sayi)
	println(sayi)

	sayilar1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pointers.Demo2(sayilar1)

	structs.Demo1()
	structs.Demo2()

	ciftSayiToplamCn := make(chan int)
	tekSayiToplamCn := make(chan int)
	go channels.CiftSayilar(ciftSayiToplamCn)
	go channels.TekSayilar(tekSayiToplamCn)
	ciftSayiToplam, tekSayiToplam := <-ciftSayiToplamCn, <-tekSayiToplamCn

	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println(carpim)

	interfaces.Demo2()

	defer_statement.Test()

	error_handling.Demo1()

	interfaces.Demo3()

	error_handling.Demo2()
	fmt.Println(error_handling.TahminEt2(101))

	string_functions.Demo1()
	string_functions.Demo2()

	restful.Demo2()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	//project.AddProduct()
	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

}
