package main

import "fmt"

func main() {
	var a, b, result int
	var success bool

	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(err)
	} else if (a == 0) && (b == 0) {
		fmt.Println("Retry input number")
		fmt.Scan(&a, &b)

		if n == 0 {
			success = false
			fmt.Println(success)
		} else {
			result = a + b
			fmt.Println(result)
			success = true
			fmt.Println(success)
		}
	}
}
