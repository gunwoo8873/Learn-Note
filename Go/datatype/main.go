package main

import (
	"fmt"
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
	fmt.Println("Integer 8bit min :", int8_min.int8_type)

	var int8_max = Integer{
		int8_type: 127,
	}
	fmt.Println("Integer 8bit max :", int8_max.int8_type)
}

// Note : Integer datatype allow value is min -128 and max 127 = 1Byte

func Integer_16() {
	var int16_min = Integer{
		int16_type: -32768,
	}
	fmt.Println("Integer 16bit min :", int16_min.int16_type)

	var int16_max = Integer{
		int16_type: 32767,
	}
	fmt.Println("Integer 16bit max :", int16_max.int16_type)
}

// Note : Integer datatype allow value is min -32768 and max 32767 = 2Byte

func Integer_32() {
	var int32_min = Integer{
		int32_type: -2147483648,
	}
	fmt.Println("Integer 16bit min :", int32_min.int32_type)

	var int32_max = Integer{
		int32_type: 2147483647,
	}
	fmt.Println("Integer 16bit max :", int32_max.int32_type)
}

// Note : Integer datatype allow value is min -2147483648[-2^31] and max 2147483647[2^31-1] = 4Byte

func Integer_64() {
	var int64_min = Integer{
		int64_type: -9223372036854775808,
	}
	fmt.Println("Integer 16bit min :", int64_min.int64_type)

	var int64_max = Integer{
		int64_type: 9223372036854775807,
	}
	fmt.Println("Integer 16bit max :", int64_max.int64_type)
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
	fmt.Println("Uinteger 8bit min :", uint8_min.uint8_type)

	var uint8_max = Uinteger{
		uint8_type: 255,
	}
	fmt.Println("Uinteger 8bit max :", uint8_max.uint8_type)
}

// Note : Uinteger datatype allow value is min 0 and max 255 = 1Byte

func Uinteger_16() {
	var uint16_min = Uinteger{
		uint16_type: 0,
	}
	fmt.Println("Uinteger 16bit min :", uint16_min.uint16_type)

	var uint16_max = Uinteger{
		uint16_type: 65535,
	}
	fmt.Println("Uinteger 16bit max :", uint16_max.uint16_type)
}

// Note : Uinteger datatype allow value is min 0 and max 65535 = 2Byte

func Uinteger_32() {
	var uint32_min = Uinteger{
		uint32_type: 0,
	}
	fmt.Println("Uinteger 32bit min :", uint32_min.uint32_type)

	var uint32_max = Uinteger{
		uint32_type: 4294967295,
	}
	fmt.Println("Uinteger 32bit max :", uint32_max.uint32_type)
}

// Note : Uinteger datatype allow value is min 0 and max 4294967295[2^32-1] = 4Byte

func Uinteger_64() {
	var uint64_min = Uinteger{
		uint64_type: 0,
	}
	fmt.Println("Uinteger 64bit min :", uint64_min.uint64_type)

	var uint64_max = Uinteger{
		uint64_type: 18446744073709551615,
	}
	fmt.Println("Uinteger 64bit max :", uint64_max.uint64_type)
}

// Note : Uinteger datatype allow value is min 0 and max 18446744073709551615[2^64-1] = 8Byte

func main() {
	Integer_8()
	Integer_16()
	Integer_32()
	Integer_64()

	Uinteger_8()
	Uinteger_16()
	Uinteger_32()
	Uinteger_64()
}
