package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	fmt.Println(strings.Count(strings.ToLower("Animatrix"), "a")) //tolower sayesinde a karakterinden iki tane buldu

	fmt.Println(strings.ToUpper("Gopher"))
}
