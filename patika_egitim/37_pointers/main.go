package main

import "fmt"

func main() {
	name := "Rıfkı"
	fmt.Println(name)
	fmt.Println(&name) // & --> address operator // başka bir değişkenin adresini tutan yerdir

	x := 22
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println("")

	fmt.Printf("%T,%v\n", x, x)
	fmt.Printf("%T,%v\n", &x, &x)

	y := &x //var y *int = &x
	fmt.Println(y)
	fmt.Printf("%T,%v\n", y, y)

	z := &name
	fmt.Printf("%T,%v\n", z, z)

	fmt.Println(x)           // 22
	fmt.Println(&x)          // 22 adresi ---> ilk *int olan veri tipidir
	fmt.Println(*(&x))       // tekrar 22 / * deadress --> pointerı tekrar adresteki değeri döndürür
	fmt.Println(&(*(&x)))    // * ----> ilgili adresteki değeri gösterir
	fmt.Println(*(&(*(&x)))) // & -->ilgili değerin adresini göterir

	fmt.Println("")
	fmt.Println(3 * 5)

	fmt.Println("")

	x1 := 10
	x2 := x1
	fmt.Println(x1, x2)
	x1 = 5
	fmt.Println(x1, x2) // sadece değeri kopyaladığı için x2 değişmez

	fmt.Println("")

	y1 := 10
	y2 := &y1 // pointer ile atama referans noktasına müdahele ettiği için veri değişir

	fmt.Println(y1, y2)
	fmt.Println(y1, *y2)
	y1 = 5
	fmt.Println(y1, *y2)
	*y2 = 3
	fmt.Println(y1, *y2)

	y3 := &y1
	*y3 = 5
	fmt.Println(y1, *y2, *y3)

	fmt.Println("")
	fmt.Println("")

	//z1 := [4]int{1, 10, 100, 1000} // array pass by value
	z1 := []int{1, 10, 100, 1000} // slice pass by reference
	z2 := z1
	fmt.Println(z1, z2)

	fmt.Println("")
	z2[0] = 3
	fmt.Println(z2)
	fmt.Println(z1)

}
