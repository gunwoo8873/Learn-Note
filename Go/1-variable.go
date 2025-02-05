package main

import "fmt"

var (
	string_type   string
	int_type_a    int
	int_type_b    uint
	float_32_type float32
	float_64_type float64
	bool_type     bool
)

func string_fn() {
	string_type = "Hello World"

	fmt.Println("string: ", string_type)
}

func integer_fn() {
	int_type_a = -15
	int_type_b = 123456789

	fmt.Println("int_a: ", int_type_a)
	fmt.Println("int_b: ", int_type_b)
}

func integer_change() {
	a := 123456789
	b := uint(a)
	c := float32(b)

	fmt.Println("c: ", c)
}

// Integers is min size 8 bit and max size 64 bit

func float_fn() {
	float_32_type = 10.10
	float_64_type = 10.10

	fmt.Println("float_32: ", float_32_type)
	fmt.Println("float_64: ", float_64_type)
}

func boolean_fn() {
	bool_type = true

	fmt.Println("boolean:  ", bool_type)
}

func main() {
	string_fn()
	integer_fn()
	integer_change()
	float_fn()
	boolean_fn()
}
