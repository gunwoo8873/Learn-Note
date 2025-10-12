package main

import (
	"fmt"
)

//// For 루프 원리
// for 초기문; 조건문; 후처리 {...}
// 초기문 생략 [O]
// 조건문 생략 [X]
// 후처리 생략 [O]

func main() {
	var loopNumber int

	fmt.Println("Enter the number of times you want of loop to run :")
	fmt.Scan(&loopNumber)

	for {
		fmt.Println("Count :", loopNumber)
		loopNumber--

		if loopNumber == 0 {
			fmt.Println("Complete loop")
			break
		}
	}
}
