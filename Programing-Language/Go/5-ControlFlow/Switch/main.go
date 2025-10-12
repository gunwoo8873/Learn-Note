package main

import "fmt"

func inputNum() int {
	var a int
	fmt.Scan(&a)
	return a
}

func main() {
	switch v := inputNum(); v {
	case 1:
		fmt.Println("+1 Value is :", v+1)
	case 2:
		fmt.Println("+2 Value is :", v+2)
	case 3:
		fmt.Println("+3 Value is :", v+3)
	default:
		fmt.Println("Value is :", v)
	}
}
