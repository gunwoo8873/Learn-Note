package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}
