package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	i := 0

	for i = 0; i < len(slice); i++ {
		slice[i] += 10 // Push <slice> index change to += 10
	}

	for i, v := range slice {
		slice[i] = v * 2 // Value * 2 calculate after to return
	}

	fmt.Println(slice)
}
