package main

import (
	"fmt"
	"time"
)

//// For 루프 원리
// for 초기문; 조건문; 후처리 {...}
// 초기문 생략 [O]
// 조건문 생략 [X]
// 후처리 생략 [O]

func standardLoop() {
	i := 0

	//// 일반적인 for 루프
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++

		if i == 5 {
			break
		}
	}

	//// 순회 루프
	for i := range 10 {
		fmt.Print(i)
	}
}

func main() {
	standardLoop()
}
