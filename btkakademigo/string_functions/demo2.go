package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {

	isim := "Umut"
	fmt.Println(s.HasPrefix(isim, "U"))
	fmt.Println(s.HasSuffix(isim, "ut"))
	fmt.Println(s.Index(isim, "ut"))
	harfler := []string{"u", "m", "u", "t"}
	fmt.Println(s.Join(harfler, "-")) // u-m-u-t
	fmt.Println(s.Join(harfler, "*")) // u*m*u*t

	fmt.Println(s.Replace(isim, "U", "Ü", 1))  // Ümut
	fmt.Println(s.Replace(isim, "U", "Ü", -1)) // Ümut

	fmt.Println(s.Split(isim, "")) // [U m u t]

	fmt.Println()
}
