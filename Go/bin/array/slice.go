package array

import "fmt"

func Slice() {
	empty := []string{}
	words := []string{"hello", "world", "golang"}

	fmt.Println("Empty Slice : %d, capacity : %d, %v", len(empty), cap(empty), empty)
	fmt.Println("Words Slice : %d, capacity : %d, %v", len(words), cap(words), words)
}
