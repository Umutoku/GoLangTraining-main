package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Umut"
	fmt.Println(s.Count(s.ToLower(isim), "u"))

}
