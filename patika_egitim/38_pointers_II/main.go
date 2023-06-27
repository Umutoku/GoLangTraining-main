package main

import "fmt"

func main() { // Go pass by value
	x := 5
	fmt.Println(x)
	double(x)
	fmt.Println(x)

	fmt.Println("")
	fmt.Println("")
	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc)
	doubleSlc(mySlc)
	fmt.Println(mySlc) // slice değişir ise underlying array değişir

	myArr := [3]int{1, 10, 100}
	fmt.Println(myArr)
	doubleArr(myArr)
	fmt.Println(myArr)

	fmt.Println("")
	fmt.Println("")

	z := 10
	fmt.Println(z)
	doubleP(&z)
	fmt.Println(z)
}

func doubleP(num *int) { //double --> pointer method
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}

func double(num int) {
	num *= 2
	fmt.Println(num)
}

func doubleSlc(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)
}

func doubleArr(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr)
}
