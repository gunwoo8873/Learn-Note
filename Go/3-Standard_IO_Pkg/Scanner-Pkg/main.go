package main

import "fmt" // 기본으로 제공하는 'fmt' 패키지

func Scan(a int) int {
	fmt.Scan(&a)

	for a == 0 {
		fmt.Println("Please enter the number again :")
		fmt.Scan(&a) // 변수들의 메모리 주소를 인수로 수신 여러 값을 입력시 변수 사이 공란을 두어 구분

		if a > 0 {
			fmt.Println("Positive number:", a)
			break
		}
	}

	return a
}

func main() {
	Scan(0)
	fmt.Println("End of program")
}
