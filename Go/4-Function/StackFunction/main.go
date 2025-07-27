package main

import "fmt"

func sumFn(a, b float32) float32 {
	return a + b
}

func minusFn(a, b float32) float32 {
	return a - b
}

func main() {
	x := sumFn(1.0, 3.1)
	fmt.Println(x)

	y := minusFn(1.0, 3.1)
	fmt.Println(y)
}
