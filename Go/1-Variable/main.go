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

	//// constant is one variable to only one value
	// v = "Error"
	// fmt.Println(v)
}

func shortKeyword() {
	//// ':=' is not use 'var' and datatype but can use short changed
	v := "This is short variable"
	fmt.Println(v)

	//// Do not equle datatype can't use
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
