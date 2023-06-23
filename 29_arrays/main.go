package main

import "fmt"

func main() {
	/* city1 := "İStanbul"
	city2 := "roma"
	city3 := "tahran"
	city4 := "floransa"
	fmt.Println(city1, city2, city3, city4) */

	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	fmt.Println(cities)
	fmt.Println(cities[0]) // zero based index
	fmt.Println(len(cities))

	var myArr [5]int // int için null value 0dır. Eğer dizi uzunluğu belli ise içine null değer ile oluşturulur.
	fmt.Println(myArr)
	myArr[0] = 100
	fmt.Println(myArr)

	var myArr1 [5]int
	var myArr2 [4]int
	fmt.Println(myArr1, myArr2)

	// if myArr1 == myArr2{ // dizinin eleman sayısı değişken tipi olarak kaydedilir.
	// 	fmt.Println("Mesaj")
	// }

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	fmt.Println("")
	myArr3 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySqaure(myArr3) // First Class Functions

	fmt.Println("")
	// FOR --- RANGE
	// for index, value := range myArr
	for index, city := range cities {
		fmt.Println(index, city)
	}

}

func mySqaure(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		fmt.Println(i, "karesi:", arr[i])
	}
	return arr
}
