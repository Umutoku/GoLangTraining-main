package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		} else {
			fmt.Println("Generic error", err)
			return
		}
	}
	fmt.Println(f.Name(), "opened successfully")
}
