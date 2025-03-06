package main

import "fmt"

func pointer(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1

	pointer(&i)
	fmt.Println("value:", i)
	fmt.Println("pointer:", &i)
}
