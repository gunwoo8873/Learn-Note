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

		/////////////////////////
		// %x = integer 16 : a-f
		// %X = integer 16 : A-F
		/////////////////////////
	)
}

func floatPrint() {
	float_value := 100.0000000

	fmt.Printf("Float Print: %f\n", float_value)
	fmt.Printf("Float Print: %g\n", float_value)
}

func stringPrint() {
	str_value := "string value"

	fmt.Printf("String Print: %s\n", str_value)
}

func memoryPrint() {
	str_value := "string value"

	fmt.Printf("Memory String Print: %p\n", &str_value)

	int_value := 100

	fmt.Printf("Memory Integer Print: %p\n", &int_value)
}

func whithPrint() {
	int_value_a := 100
	int_value_b := 150000.00000

	fmt.Printf("Output whith size [5]Point Print: %5d, %5g\n", int_value_a, int_value_b)
	fmt.Printf("Output whith left align [5]Point Print: %-5d, %-5g\n", int_value_a, int_value_b)
	fmt.Printf("Output whith min size add 0 [5]Point Print: %05d, %05g\n", int_value_a, int_value_b)
	fmt.Printf("Output Float [5]langth %8.3f\n", int_value_b)

	//////////////////////////////////////////////////
	// %5d 		= 5 digit,
	// %-5d	 	= 5 digit left align,
	// %05d 	= 5 digit min size add 0 [00100, 150000]
	// %N.Nf 	= [N] digit, [N] digit after point
	//////////////////////////////////////////////////
}

func main() {
	standard()
	dataTypePrint()
	booleanPrint()
	integerPrint()
	floatPrint()
	stringPrint()
	memoryPrint()
	whithPrint()
}
