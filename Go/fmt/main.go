package main

import "fmt"

func standard() {
	var int_value = 100
	fmt.Printf("Standard Print: %v\n", int_value)
}

func dataTypePrint() {
	int_value := 100
	str_value := "string value"
	fmt.Printf("Data type Print: %T, %T\n", int_value, str_value)
}

func booleanPrint() {
	boolean_value := true
	fmt.Printf("Boolean Print: %t\n", boolean_value)
}

func integerPrint() {
	int_16_value := 10 // 16진수로 출력할 값
	int_10_value := 10 // 10진수로 출력할 값
	int_8_value := 10  // 8진수로 출력할 값
	int_2_value := 10  // 2진수로 출력할 값
	unicode_value := '1'

	fmt.Printf(
		"Integer 16 Print: %x\nInteger 10 Print: %d\nInteger 8 Print: %o\nInteger 2 Print: %b\nUnicode Print: %c\n",
		int_16_value, int_10_value, int_8_value, int_2_value, unicode_value,
		// %x = integer 16 : a-f, %X = integer 16 : A-F
	)
}

func mathPrint() {
	float_value := 100.0
	fmt.Printf("Math Print: %f\n", float_value)
}

func main() {
	standard()
	dataTypePrint()
	booleanPrint()
	integerPrint()
	mathPrint()
}
