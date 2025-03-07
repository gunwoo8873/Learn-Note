package main

import "fmt"

func main() {
	empty := []int{}
	fmt.Printf("empty: length: %d, capacity: %d, %v\n", len(empty), cap(empty), empty)

	words := []int{0, 1, 2, 3, -1, -2, -3}
	fmt.Printf("words: length: %d, capacity: %d, %v\n", len(words), cap(words), words)

	add := make([]int, 4)
	fmt.Printf("add: length: %d, capacity: %d, %v\n", len(add), cap(add), add)

	psuh := append(add, 5) // append = push or insert
	fmt.Printf("push: length: %d, capacity: %d, %v\n", len(psuh), cap(psuh), psuh)
}
