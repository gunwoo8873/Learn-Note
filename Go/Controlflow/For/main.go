package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func main() {
	var a = make([]int, 30) // 30만큼 Slice 생성

	// rnage : 데이터 타입을 순회할 때 사용되는 키워드 반환값은 Index와 Value를 반환
	for i := range a {
		a[i] = i + 1 // Slice의 배열에 값을 i + 1만큼 반복 저장
	}

	fmt.Println(a)

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}
