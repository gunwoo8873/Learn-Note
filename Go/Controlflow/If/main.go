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

func multipleIFLoop() {
	a := make([]int, 10)

	for i := range a {
		a[i] = i + 1

		if a[i] == 10 {
			if a[i] == 3 {
				fmt.Println(a[i])
				continue
			}

			if a[i] == 6 {
				fmt.Println(a[i])
				continue
			}

			if a[i] == 9 {
				fmt.Println(a[i])
				break
			}
		}
	}
}

func main() {
	arrLoop()
	multipleIFLoop()
}
