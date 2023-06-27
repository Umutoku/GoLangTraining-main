package main

import "fmt"

func main() {
	myArr := []int{1, 2, 3}
	myArr1 := [...]int{1, 2, 3, 4}
	fmt.Println(myArr)
	fmt.Println(len(myArr))
	fmt.Println(myArr1)
	fmt.Println(len(myArr1))

	var myArr3 [5]int
	fmt.Println(myArr3)

	mySlc := []int{1, 2, 3} //literal yöntemi
	fmt.Println(mySlc)
	fmt.Println(len(mySlc))

	var myArr4 [4]int
	fmt.Println(myArr4)

	var mySlc1 []int // make method
	mySlc1 = make([]int, 4)
	fmt.Println(mySlc1)

	fmt.Println("")
	fmt.Println("")

	myArrEx := [3]int{1, 2, 3}
	fmt.Println(myArrEx)
	myArrEx1 := myArrEx
	fmt.Println(myArrEx1)
	myArrEx1[0] = 100 // arraylerde diziler birbirine eşitlenirse diğer arraylerde değişiklik yapmak referans alınan array'i etkilemez
	fmt.Println(myArrEx1)
	fmt.Println(myArrEx) // pass by value

	mySlcEx := []int{1, 2, 3}
	fmt.Println(mySlcEx)
	mySlcEx1 := mySlcEx
	fmt.Println(mySlcEx1)
	mySlcEx1[0] = 33
	fmt.Println(mySlcEx1)
	fmt.Println(mySlcEx) //pass by referance

}
