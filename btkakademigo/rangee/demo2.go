package rangee

import "fmt"

func Demo2() {
	var total int = 0
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		if i%2 != 0 {
			total += i
		}
	}
	fmt.Println(total)
}
