package defer_statement

import "fmt"

type product struct {
	name  string
	price int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi", p.name)
	defer Log()
}

func Log() {
	fmt.Println("loglandı")
}

func Demo3() {
	p := product{name: "Laptop", price: 5000}
	defer p.save()
	fmt.Println("İşlem başarılı")
}
