package main

import "fmt"

func main() {
	myMap := map[string]int{
		"Ahmet":   40,
		"Ayşe":    30,
		"Hayriye": 25,
	}

	fmt.Println(myMap)
	//fmt.Println(myMap["Ahmet"])
	//fmt.Println(myMap["Ahmet"], myMap["Ayşe"]) ------ //map içinde iki veri çağıracak isek map adını tanımlamalıyız

	fmt.Println(myMap["Olmayan"]) //eğer map içerisinde yoksa string nil int 0 olarak dönüş yapar.

	isMarride := map[string]bool{
		"Metin": false,
		"Berke": false,
		"Pelin": false,
	}
	fmt.Println(isMarride)

	myMap1 := make(map[string]int)
	fmt.Println(myMap1)

	studentGrades := map[string]int{
		"Ali":   80,
		"Veli":  49,
		"Fevzi": 50,
		"Çınar": 0,
	}
	fmt.Println(studentGrades["Ali"])
	fmt.Println(studentGrades["Çınar"]) //notu 0 olduğu için olmayan bir değer ile aynı çıkar

	value, ok := studentGrades["Demet"] //ok kontrolü false dönüyor ise bu değer bulunmamakta
	fmt.Println(value, ok)

	fmt.Println(studentGrades)
	studentGrades["Mahmut"] = 55
	fmt.Println(studentGrades) //üst satır map'e eklenecek

	delete(studentGrades, "Fevzi") //eleman silmek için delete methodu kullanılabilir
	fmt.Println(studentGrades)

	fmt.Println(len(studentGrades))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	sg := studentGrades //maps --> pass by reference // olan yeri paylaşır

	fmt.Println(sg)
	fmt.Println(studentGrades)
	delete(sg, "Mahmut") //maplerde slice gibi referansı yani aynı değeri siliyor. Array gibi veriyi kendi saklamıyor
	fmt.Println(sg)
	fmt.Println(studentGrades)

	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
}
