package main

import "fmt"

func arithmetic() {
	var a, b int = 15, 40

	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Printf(
		"Sum: %d\nDifference: %d\nProduct: %d\nQuotient: %d\nRemainder: %d\n",
		sum, diff, product, quotient, remainder,
	)
}

func main() {
	arithmetic()
}
