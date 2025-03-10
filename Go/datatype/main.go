package main

import (
	"fmt"
	"math"
)

// golan is integer and uinteger do have't 128bit type
type Integer struct {
	int8_type  int8
	int16_type int16
	int32_type int32
	int64_type int64
}

func Integer_8() {
	var int8_min = Integer{
		int8_type: -128,
	}

	var int8_max = Integer{
		int8_type: 127,
	}

	var int8_min_math int8 = math.MinInt8
	var int8_max_math int8 = math.MaxInt8

	if int8_min_math == int8_min.int8_type && int8_max_math == int8_max.int8_type {
		fmt.Printf("min int 8: %v = %v, max int 8: %v = %v\n", int8_min_math, int8_min.int8_type, int8_max_math, int8_max.int8_type)
	} else {
		fmt.Printf("min int 8: %v = %v, max int 8: %v = %v\n", int8_min_math, int8_min.int8_type, int8_max_math, int8_max.int8_type)
	}
}

// Note : Integer datatype allow value is min -128 and max 127 = 1Byte

func Integer_16() {
	var int16_min = Integer{
		int16_type: -32768,
	}

	var int16_max = Integer{
		int16_type: 32767,
	}

	var int16_min_math int16 = math.MinInt16
	var int16_max_math int16 = math.MaxInt16

	if int16_min_math == int16_min.int16_type && int16_max_math == int16_max.int16_type {
		fmt.Printf("min int 16: %v = %v, max int 16: %v = %v\n", int16_min_math, int16_min.int16_type, int16_max_math, int16_max.int16_type)
	} else {
		fmt.Printf("min int 16: %v = %v, max int 16: %v = %v\n", int16_min_math, int16_min.int16_type, int16_max_math, int16_max.int16_type)
	}
}

// Note : Integer datatype allow value is min -32768 and max 32767 = 2Byte

func Integer_32() {
	var int32_min = Integer{
		int32_type: -2147483648,
	}

	var int32_max = Integer{
		int32_type: 2147483647,
	}

	var int32_min_math int32 = math.MinInt32
	var int32_max_math int32 = math.MaxInt32

	if int32_min_math == int32_min.int32_type && int32_min_math == int32_max.int32_type {
		fmt.Printf("min int 32: %v = %v, max int 32: %v = %v\n", int32_min_math, int32_min.int32_type, int32_max_math, int32_max.int32_type)
	} else {
		fmt.Printf("min int 32: %v = %v, max int 32: %v = %v\n", int32_min_math, int32_min.int32_type, int32_max_math, int32_max.int32_type)
	}
}

// Note : Integer datatype allow value is min -2147483648[-2^31] and max 2147483647[2^31-1] = 4Byte

func Integer_64() {
	var int64_min = Integer{
		int64_type: -9223372036854775808,
	}

	var int64_max = Integer{
		int64_type: 9223372036854775807,
	}

	var int64_min_math int64 = math.MinInt64
	var int64_max_math int64 = math.MaxInt64

	if int64_min_math == int64_min.int64_type && int64_max_math == int64_max.int64_type {
		fmt.Printf("min int 64: %v = %v, max int 64: %v = %v\n", int64_min_math, int64_min.int64_type, int64_max_math, int64_max.int64_type)
	} else {
		fmt.Printf("min int 64: %v = %v, max int 64: %v = %v\n", int64_min_math, int64_min.int64_type, int64_max_math, int64_max.int64_type)
	}
}

// Note : Integer datatype allow value is min -9223372036854775808[-2^64] and max 9223372036854775807[2^64-1] = 8Byte

type Uinteger struct {
	uint8_type  uint8
	uint16_type uint16
	uint32_type uint32
	uint64_type uint64
}

func Uinteger_8() {
	var uint8_min = Uinteger{
		uint8_type: 0,
	}

	var uint8_max = Uinteger{
		uint8_type: 255,
	}

	var uint8_min_math uint8 = 0 // We can just uint 8bit is 0 directly
	var uint8_max_math uint8 = math.MaxUint8

	if uint8_min_math == uint8_min.uint8_type && uint8_max_math == uint8_max.uint8_type {
		fmt.Printf("min uint 8: %v = %v, max uint 8: %v = %v\n", uint8_min_math, uint8_min.uint8_type, uint8_max_math, uint8_max.uint8_type)
	} else {
		fmt.Printf("min uint 8: %v = %v, max uint 8: %v = %v\n", uint8_min_math, uint8_min.uint8_type, uint8_max_math, uint8_max.uint8_type)
	}
}

// Note : Uinteger datatype allow value is min 0 and max 255 = 1Byte

func Uinteger_16() {
	var uint16_min = Uinteger{
		uint16_type: 0,
	}

	var uint16_max = Uinteger{
		uint16_type: 65535,
	}

	var uint16_min_math uint16 = 0 // We can just uint 8bit is 0 directly
	var uint16_max_math uint16 = math.MaxUint16

	if uint16_min_math == uint16_min.uint16_type && uint16_max_math == uint16_max.uint16_type {
		fmt.Printf("min uint 16: %v = %v, max uint 16: %v = %v\n", uint16_min_math, uint16_min.uint16_type, uint16_max_math, uint16_max.uint16_type)
	} else {
		fmt.Printf("min uint 16: %v = %v, max uint 16: %v = %v\n", uint16_min_math, uint16_min.uint16_type, uint16_max_math, uint16_max.uint16_type)
	}
}

// Note : Uinteger datatype allow value is min 0 and max 65535 = 2Byte

func Uinteger_32() {
	var uint32_min = Uinteger{
		uint32_type: 0,
	}

	var uint32_max = Uinteger{
		uint32_type: 4294967295,
	}

	var uint32_min_math uint32 = 0 // We can just uint 8bit is 0 directly
	var uint32_max_math uint32 = math.MaxUint32

	if uint32_min_math == uint32_min.uint32_type && uint32_max_math == uint32_max.uint32_type {
		fmt.Printf("min uint 32: %v = %v, max uint 32: %v = %v\n", uint32_min_math, uint32_min.uint32_type, uint32_max_math, uint32_max.uint32_type)
	} else {
		fmt.Printf("min uint 32: %v = %v, max uint 32: %v = %v\n", uint32_min_math, uint32_min.uint32_type, uint32_max_math, uint32_max.uint32_type)
	}
}

// Note : Uinteger datatype allow value is min 0 and max 4294967295[2^32-1] = 4Byte

func Uinteger_64() {
	var uint64_min = Uinteger{
		uint64_type: 0,
	}

	var uint64_max = Uinteger{
		uint64_type: 18446744073709551615,
	}

	var uint64_min_math uint64 = 0 // We can just uint 8bit is 0 directly
	var uint64_max_math uint64 = math.MaxUint64

	if uint64_min_math == uint64_min.uint64_type && uint64_max_math == uint64_max.uint64_type {
		fmt.Printf("min uint 64: %v = %v, max uint 64: %v = %v\n", uint64_min_math, uint64_min.uint64_type, uint64_max_math, uint64_max.uint64_type)
	} else {
		fmt.Printf("min uint 64: %v = %v, max uint 64: %v = %v\n", uint64_min_math, uint64_min.uint64_type, uint64_max_math, uint64_max.uint64_type)
	}
}

// Note : Uinteger datatype allow value is min 0 and max 18446744073709551615[2^64-1] = 8Byte

type Float struct {
	float32_type float32
	float64_type float64
}

func Float_32() {
	var float_32 = Float{
		float32_type: 0.123456789,
	}
	fmt.Println("float 32 :", float_32.float32_type)
}

func Float_64() {
	var float_64 = Float{
		float64_type: 0.123456789,
	}
	fmt.Println("float 64 :", float_64.float64_type)
}

// That float is digit 0.1234567 however digit 7+ last digit override

func main() {
	Integer_8()
	Integer_16()
	Integer_32()
	Integer_64()

	Uinteger_8()
	Uinteger_16()
	Uinteger_32()
	Uinteger_64()

	Float_32()
	Float_64()
}
