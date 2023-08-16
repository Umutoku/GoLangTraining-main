package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50
	if tahmin < 0 || tahmin > 100 {
		return "", errors.New("Tahmin 0 ile 100 arasında olmalıdır")
	}

	if tahmin > aklimdakiSayi {
		return "Tahmin büyük", nil
	}

	if tahmin < aklimdakiSayi {
		return "Tahmin küçük", nil
	}

	return "Tahmin doğru", nil

}

func Demo2() {
	mesaj, _ := TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(TahminEt(500))
	fmt.Println(TahminEt(50))
}
