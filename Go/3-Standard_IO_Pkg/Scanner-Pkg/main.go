package main

import "fmt" // 기본으로 제공하는 'fmt' 패키지

//// Scan : 변수들의 메모리 주소를 인수로 수신 여러 값을 입력시 변수 사이 공란을 두어 구분
func Scan(a int) int {
	fmt.Scan(&a)

	for a == 0 {
		fmt.Println("Please enter the number again :")
		fmt.Scan(&a)

		if a > 0 {
			fmt.Println("Positive number:", a)
			break
		}
	}

	return a
}

//// Scanf : 입력에서 서식 형태로 값을 읽는다
func Scanf() {
	var v int
	fmt.Println("Enter a Scanf value :")
	fmt.Scanf("%d", &v)
	fmt.Println("Scanf value:", v)
}

func Scanln() {
	var v int
	fmt.Println("Enter a Scanln value :")
	fmt.Scanln(&v)
	fmt.Println("Scanln value:", v)
}

func main() {
	Scan(0)
	Scanf()
	fmt.Println("End of program")
}
