package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func main() {
	var a = make([]int, 30)

	for i := range a {
		a[i] = i + 1

		if a[i] > 20 {
			fmt.Println(a) // A Slice value is limited to 20 [0..=20]
			fmt.Println("Slice value is over 20 and process is exit")
			break
		}
	}

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}
