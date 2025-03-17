package main

import "fmt"

func multiValue(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	a, b := multiValue(10, 3)
	fmt.Println(a, b)
}
