package main

import "fmt"

func main() {
	var x int // x 값을 초기화
	var y int // y 값을 초기화

	n, err := fmt.Scan(&x, &y) // Scan을 사용하여 각 x, y 입력값을 수신
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(x, y)
	}
}
