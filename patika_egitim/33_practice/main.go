package main

import "fmt"

func main() {
	soru4()
	soru3()
	soru2()
	soru1()
}

func soru1() {
	myArr := [5]string{
		"İstanbul",
		"Roma",
		"Paris",
		"Budapeşte",
		"Antalya",
	}
	for index, value := range myArr {
		fmt.Println(index, value)
	}

	mySlc := []int{1, 2, 3, 4, 5}
	for i, v := range mySlc {
		fmt.Println(i, "", v)
	}
}

func soru2() {
	mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	myslc1 := mySlc[:4]
	myslc2 := mySlc[4:7]
	myslc3 := mySlc[6:]
	myslc4 := append(mySlc[2:4], mySlc[6:8]...)
	fmt.Println(mySlc)
	fmt.Println(myslc1)
	fmt.Println(myslc2)
	fmt.Println(myslc3)
	fmt.Println(myslc4)

}

func soru3() {
	mySlc := []int{1, 2, 3}
	mySlc1 := make([]int, 2)
	fmt.Println(mySlc)
	fmt.Println(mySlc1)
	//mySlc1 = mySlc eğer assign edersek yani eşittir ile atarsak aynı yeri refere eder, değişiklikler aktarılır
	copy(mySlc1, mySlc) //slice kopyalaması sonrası sadece 1,2 değerlerini alır, çünkü make ile yapılan slice 2 değer tutmakta
	fmt.Println(mySlc1)

}
func soru4() {
	myMap := map[string][]string{
		"Yaşar Kemal":     []string{"İnce memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": []string{"1Q84", "Dans Dans Dans", "Kumandanı Öldürmek"},
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Sabahattin Ali"])
	fmt.Println(myMap["Yaşar Kemal"][0])

	for yazar, eser := range myMap {
		fmt.Println("Yazarımız: ", yazar)
		fmt.Println("Bazı Kitapları: ")
		for i, v := range eser {
			fmt.Println("\t", i+1, v)
		}

	}
}
