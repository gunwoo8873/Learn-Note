package main

import "fmt" // 기본으로 제공하는 'fmt' 패키지

func datatype() {
	a := 15

	fmt.Printf("%v \n", a) // 값의 데이터 타입을 기본 형태로 출력
	fmt.Printf("%T \n", a) // 값의 데이터 타입을 출력
}

func integer() {
	a := 15

	fmt.Printf("%d \n", a) // 10진수 정숫값으로 출력
	fmt.Printf("%b \n", a) // 2진수로 출력
	fmt.Printf("%o \n", a) // 8진수로 출력
	fmt.Printf("%O \n", a) // 'Oo'포함하여 8진수로 출력
	fmt.Printf("%x \n", a) // 16진수로 출력 10 이상 값은 'a-f'
	fmt.Printf("%X \n", a) // 대문자 16진수로 출력 10 이상 값은 'A-F'
	fmt.Printf("%U \n", a) // 유니코드 문자로 출력
}

func float() {
	a := 1.23456789

	fmt.Printf("%e \n", a) // 지수 형태 실숫값 출력
	fmt.Printf("%E \n", a)

	fmt.Printf("%f \n", a) // 값의 실숫값 출력
	fmt.Printf("%F \n", a)

	fmt.Printf("%g \n", a) // 값이 큰 실숫값은 지수 형태로 출력 '%e'
	fmt.Printf("%G \n", a) // 값이 작은 실숫값은 지수 형태로 출력 '%f'
}

func boolean() {
	a := true

	fmt.Printf("%t \n", a) // Boolean의 true / false 값으로 출력
}

func string() {
	a := "String"

	fmt.Printf("%s \n", a) // 문자열 출력
	fmt.Printf("%q \n", a) // 특수 문자 기능을 동작하지 않고 모든 문자열 출력
}

func memoryAddress() {
	a := 10

	fmt.Printf("%p \n", &a) // 메모리 주솟값을 출력
}

func main() {
	fmt.Println("") // 함수 입력값 출력 및 개행
	fmt.Printf("")  // 'Format'에 맞게 입력값 출력
	fmt.Print("")   // 함수 입력값 출력

	fmt.Scan("")   // 표준 입력에서 값을 입력받는다
	fmt.Scanf("")  // 표준 입력에서 서식 형태로 값을 입력받는다
	fmt.Scanln("") // 표준 입력에서 한 줄을 읽어서 값을 입력받는다

	datatype()
	boolean()
	integer()
	float()
	string()
	memoryAddress()
}
