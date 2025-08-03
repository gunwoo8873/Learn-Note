package main

import "fmt"

func varKeyowrd() {
	var v string = "This is variable"
	fmt.Println(v)

	v = "This is update after value"
	fmt.Println(v)
}

func constKeyword() {
	const v string = "This is constant"
	fmt.Println(v)

	//// [X] const는 하나의 값만 활당이 가능하며, 값의 변경이 불가능하다.
	// v = "Error"
	// fmt.Println(v)
}

func shortKeyword() {
	//// ':='는 'var'과 같은 기능을 하며, 데이터 타입에 대한 명시를 생략을 할 수 있다. [단. 변수는 반드시 선언이 되어야 하며, 값이 없는 초기화는 불가능하다.]
	v := "This is short variable"
	fmt.Println(v)

	//// [X] 초기 선언과의 데이터 타입이 다르기 때문에, 재활당이 불가능하다.
	// v := 10
	// fmt.Println(v)

	v = "This is update after short variable"
	fmt.Println(v)
}

func main() {
	varKeyowrd()
	constKeyword()
	shortKeyword()
}
