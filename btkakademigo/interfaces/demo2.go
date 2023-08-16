package interfaces

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {

	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {

	return c.creditPaymentTotal * c.rate / 100 / 12
}

func MonthlyPayment(credits []CreditCalculator) float64 {

	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal += credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{100000, "address1", 10}
	credit3 := Mortgage{40000, "address1", 15}
	credit2 := Car{20000, "car1", 12}
	credits := MonthlyPayment([]CreditCalculator{credit1, credit2, credit3})

	println(credits)
}
