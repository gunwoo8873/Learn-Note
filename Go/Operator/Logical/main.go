package main

import "fmt"

func logical() {
	var a, b int = 4, 8

	and := a & b // OR a && b 1 & 1 = 1, 0 & 1 = 0
	or := a | b  // OR a || b 1 | 1 = 1, 0 | 1 = 1, 0 | 0 = 0
	xor := a ^ b // 1 ^ 1 = 0, 0 ^ 1 = 1, 0 ^ 0 = 0
	not := ^a    // OR !a

	fmt.Printf(
		"And: %d, Or: %d, Xor: %d, Not: %d\n",
		and, or, xor, not,
	)

	equal := a == b
	notEqual := a != b
	lessThan := a < b
	lessThanEqual := a <= b
	greaterThan := a > b
	greaterThanEqual := a >= b

	fmt.Printf(
		"Equal: %t\nNot Equal: %t\nLess Than: %t\nLess Than Equal: %t\nGreater Than: %t\nGreater Than Equal: %t\n",
		equal, notEqual, lessThan, lessThanEqual, greaterThan, greaterThanEqual,
	)

	///////////////////////////////////////
	// || 										[1순위 연산자]
	// && 										[2순위 연산자]
	// ==, !=, <, <=, >, >= 	[3순위 연산자]
	// +, -, |, ^ 						[4순위 연산자]
	// *, /, %, <<, >>, &, &^ [5순위 연산자]
	///////////////////////////////////////
}

func stringLogical() {
	stra := "var"
	strb := "const"
	strc := "var"

	fmt.Printf("%s == %s : %v\n", stra, strb, stra == strb)
	fmt.Printf("%s != %s : %v\n", stra, strb, stra != strb)
	fmt.Printf("%s == %s : %v\n", stra, strc, stra == strc)
	fmt.Printf("%s != %s : %v\n", stra, strc, stra != strc)
}

func main() {
	logical()
	stringLogical()
}
