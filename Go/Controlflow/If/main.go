package main

import "fmt"

func arrLoop() {
	a := make([]int, 10)

	for i := range a {
		a[i] = i + 1

		if a[i] > 5 {
			fmt.Println(a) // A Slice value is limited to 20 [0..=20]
			fmt.Println("Slice value is over 20 and process is exit")
			break
		}
	}
}

func main() {
	arrLoop()
}
