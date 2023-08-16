package rangee

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "table": "masa", "computer": "bilgisayar"}

	for key, value := range sozluk {
		println(key, value)
	}
}
