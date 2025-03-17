package main

import "fmt"

func recursive(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	recursive(n - 1)
	fmt.Println(n)
}

func main() {
	recursive(3)
}
