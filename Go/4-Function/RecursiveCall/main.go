package main

import "fmt"

func Recursive(x int) {
	if x == 0 {
		return
	}

	fmt.Println("Befor Result :", x)
	Recursive(x - 1)
	fmt.Println("After Result :", x)
}

func main() {
	Recursive(5)
}
