package main

import "fmt"

func main() {
	const x int16 = 5
	fmt.Printf("%T,%v", x, x)

	const y = 3 //typeless veri
	var z int16 = 12
	fmt.Printf("%T,%v", y, y)
	fmt.Printf("%T,%v", z, z)
	fmt.Printf("%T,%v", (y + z), (y + z)) // int16(y)+z // type conversion kullanıldığında veri tipi olur

	const k = 5.2 + 4.8
	fmt.Printf("%T,%v", k, k)
}
