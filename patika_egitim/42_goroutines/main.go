package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() { //main goroutine main çalışmaya başladığı zaman başlıyor

	wg.Add(1)

	go printX()
	wg.Wait()
	fmt.Println()
	printY()
	time.Sleep(time.Second)
}

func printX() {
	for i := 0; i < 20; i++ {
		fmt.Println("X")
	}
	wg.Done()
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Println("Y")
	}
}
