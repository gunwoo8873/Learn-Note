package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func pointer(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1

	pointer(&i)
	fmt.Println("value:", i)
	fmt.Println("pointer:", &i)

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}
