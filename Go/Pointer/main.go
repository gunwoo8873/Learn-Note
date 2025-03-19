package main

import "fmt"

///////////////////////////////////////////////////////
// Pointer
// 1. Memory의 주소를 값으로 갖는 타입
// 2. 메모리 복사를 줄여 반환값 없이 변숫값을 변경
// 3. Pointer는 변숫값을 초기화하지 않으면 기본값은 nil
///////////////////////////////////////////////////////

func standardPointer() {
	var a int = 500
	var p *int // Pointer variable declaration

	fmt.Printf("Value of a: %d, Address : %p\n", a, &a)

	p = &a   // Clone address of a into p
	*p = 100 // Change value of a

	fmt.Printf("Value of p: %d, Address : %p\n", *p, p)
}

func main() {
	standardPointer()
}
