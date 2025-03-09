package main

import (
	"fmt"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Vaiabic은 모든 추가 인수는 같은 타입이어야 한다
// 추가 인수는 항상 함수의 마지막 배열의 인수여야 한다
////////////////////////////////////////////////////////////////////////////////

var start = time.Now()

func addObject(object ...string) {
	fmt.Println(strings.Join(object, "."))
}

func main() {
	addObject("127", "0", "0", "1")

	a := []string{"192", "168", "0", "1"}
	addObject(a...)

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}
