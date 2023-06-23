package main

import "fmt"

func main() {
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	mySlc := underArray[2:5]
	fmt.Println(mySlc)
	mySlc1 := underArray[:6]
	fmt.Println(mySlc1)
	mySlc2 := underArray[3:]
	fmt.Println(mySlc2)
	mySlc3 := underArray[:]
	fmt.Println(mySlc3)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println(mySlc)
	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc1)
	fmt.Println(mySlc2)
	fmt.Println(mySlc3)

	fmt.Println("")
	fmt.Println("")

	mySlcEx := []int{1, 2, 3} //mySliceUnderArray
	fmt.Println(mySlcEx)
	mySlcEx = append(mySlcEx, 4, 5)
	fmt.Println(mySlcEx)

	mySlcEx2 := append(mySlcEx, 4, 5) //mySlice2UnderArray oluşuyor
	fmt.Println(mySlcEx2)

	fmt.Println("")
	fmt.Println("")
	mySlcEx[0] = 100
	fmt.Println(mySlcEx)
	fmt.Println(mySlcEx2)
	fmt.Println("")
	fmt.Println("")

	mySlcOrn := []int{1, 2, 3}
	mySlcOrn1 := []int{4, 5}
	mySlcOrn = append(mySlcOrn, mySlcOrn1...)
	fmt.Println(mySlcOrn)

	mySlcD := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlcD)
	mySlcD = mySlcD[2:] // ilk iki eleman silinmiş olacak
	fmt.Println(mySlcD)
	mySlcD = mySlcD[:len(mySlcD)-3] // son üç elemanı silecek
	fmt.Println(mySlcD)

	var mySlcM []int
	mySlcM = make([]int, 4) // Zero değerler slice elemanlarına ait
	fmt.Println(mySlcM)

	fmt.Println("")
	fmt.Println("")
	var mySlcM1 []int // yok slice
	fmt.Printf("%#v", mySlcM1)
	fmt.Println("")

	mySlc4 := make([]int, 0) //boş slice
	fmt.Printf("%#v", mySlc4)

}
