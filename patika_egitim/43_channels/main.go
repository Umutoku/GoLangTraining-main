package main

import (
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A circle")
	wg.Done()
}

// func (c circle) area() float64 {
// 	return math.Pi * c.r * c.r // returnlerde gogoutines oluşamaz

// }

func area(c circle, myChannel chan float64) {
	myChannel <- math.Pi * c.r * c.r
}

func main() {

	//var myChannel chan string
	//myChannel = make(chan string)

	myChannel := make(chan string)

	go merhaba(myChannel)

	fmt.Println(<-myChannel)

	wg.Add(1)

	c1 := circle{5}
	chan1 := make(chan float64)
	go area(c1, chan1)
	go c1.display()
	wg.Wait()
	fmt.Printf("%.2f\n", <-chan1)

}

func merhaba(chan1 chan string) {
	chan1 <- "Merhaba"
}

// package main

// import "fmt"

// func main() {
// 	chan1 := make(chan string)

// 	chan1 <- "Arı"

// 	fmt.Println(<-chan1)
// }
