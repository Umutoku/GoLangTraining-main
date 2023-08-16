package maps

import "fmt"

func Demo1() {
	//Key-Value
	sozluk := make(map[string]string)
	sozluk["book"] = "kitap"
	sozluk["table"] = "masa"
	sozluk["computer"] = "bilgisayar"
	keys := make([]string, 0, len(sozluk))
	for k := range sozluk {
		keys = append(keys, k)
	}

	for _, k := range keys {
		fmt.Println(k, sozluk[k])
	}

	deger, varMi := sozluk["book"]
	fmt.Println(deger, varMi)

	sozluk2 := map[string]string{"glass": "bardak", "pen": "kalem"}
	for index, karsideger := range sozluk2 {
		fmt.Println("İndex ve değer: ", index, karsideger)
	}
}
