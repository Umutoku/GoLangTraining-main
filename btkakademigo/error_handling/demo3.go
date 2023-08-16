package error_handling

import "fmt"

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("parameter: %d, message: %s", b.parameter, b.message)
}

func TahminEt2(sayi int) (string, error) {
	if sayi < 1 || sayi > 100 {
		return "", &borderException{sayi, "1 ile 100 arasında bir sayı giriniz"}
	}

	return "Tahmininiz doğru", nil
}
